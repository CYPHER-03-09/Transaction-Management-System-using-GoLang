package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/rs/cors"

	"newpostgres/models"
	"newpostgres/storage"
	pb "newpostgres/pb/pb"

	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedItemServiceServer
	DB *gorm.DB
}

func (s *server) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	item := models.Item{
		ItemName:    req.ItemName,
		Description: req.Description,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add item: %v", err)
	}

	return &pb.AddItemResponse{Message: "Item added successfully"}, nil
}

func (s *server) GetItems(ctx context.Context, _ *pb.Empty) (*pb.GetItemsResponse, error) {
	items := []models.Item{}
	if err := s.DB.Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve items: %v", err)
	}

	res := &pb.GetItemsResponse{}
	for _, item := range items {
		res.Items = append(res.Items, &pb.Item{
			ItemId:      uint32(item.ItemID),
			ItemName:    item.ItemName,
			Description: item.Description,
		})
	}

	return res, nil
}

func (s *server) AddClient(ctx context.Context, req *pb.AddClientRequest) (*pb.AddClientResponse, error) {
	client := models.Client{
		ClientName: req.ClientName,
		Phone:      req.Phone,
		Email:      req.Email,
	}

	if err := s.DB.Create(&client).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add client: %v", err)
	}

	return &pb.AddClientResponse{Message: "Client added successfully"}, nil
}

func (s *server) GetClients(ctx context.Context, _ *pb.Empty) (*pb.GetClientsResponse, error) {
	clients := []models.Client{}
	if err := s.DB.Find(&clients).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve clients: %v", err)
	}

	res := &pb.GetClientsResponse{}
	for _, client := range clients {
		res.Clients = append(res.Clients, &pb.Client{
			ClientId:   uint32(client.ClientID),
			ClientName: client.ClientName,
			Phone:      client.Phone,
			Email:      client.Email,
		})
	}

	return res, nil
}

func (s *server) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest) (*pb.AddTransactionResponse, error) {
	transaction := models.Transaction{
		ClientID:   uint(req.ClientId),
		ItemID:     uint(req.ItemId),
	}

	if err := s.DB.Create(&transaction).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add transaction: %v", err)
	}

	return &pb.AddTransactionResponse{Message: "Transaction added successfully"}, nil
}

func (s *server) GetTransactions(ctx context.Context, _ *pb.Empty) (*pb.GetTransactionsResponse, error) {
	transactions := []models.Transaction{}
	if err := s.DB.Find(&transactions).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve transactions: %v", err)
	}

	res := &pb.GetTransactionsResponse{}
	for _, transaction := range transactions {
		res.Transactions = append(res.Transactions, &pb.Transaction{
			TransactionId: uint32(transaction.TransactionID),
			ClientId:      uint32(transaction.ClientID),
			//ClientName:    transaction.ClientName,
			ItemId:        uint32(transaction.ItemID),
		})
	}

	return res, nil
}

func (s *server) GetTransactionDetails(ctx context.Context, req *pb.GetTransactionDetailsRequest) (*pb.GetTransactionDetailsResponse, error) {
	type Result struct {
		TransactionID   uint
		ClientID        uint
		ClientName      string
		Phone           string
		Email           string
		ItemID          uint
		ItemName        string
		ItemDescription string
	}

	var result Result

	// Perform a JOIN query to fetch transaction details and associated item/client details
	if err := s.DB.Table("transactions").
		Select("transactions.transaction_id, transactions.client_id, clients.client_name, clients.phone, clients.email, transactions.item_id, items.item_name, items.description as item_description").
		Joins("JOIN clients ON transactions.client_id = clients.client_id").
		Joins("JOIN items ON transactions.item_id = items.item_id").
		Where("transactions.transaction_id = ?", req.TransactionId).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Transaction not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to retrieve transaction details: %v", err)
	}

	return &pb.GetTransactionDetailsResponse{
		TransactionId:   uint32(result.TransactionID),
		ClientId:        uint32(result.ClientID),
		ClientName:      result.ClientName,
		Phone:           result.Phone,
		Email:           result.Email,
		ItemId:          uint32(result.ItemID),
		ItemName:        result.ItemName,
		ItemDescription: result.ItemDescription,
	}, nil
}

// func (s *server) GetTransactionDetails(ctx context.Context, req *pb.GetTransactionDetailsRequest) (*pb.GetTransactionDetailsResponse, error) {
//     type Result struct {
//         TransactionID   uint
//         ClientID        string
//         ClientName      string
//         ItemID          uint
//         ItemName        string
//         ItemDescription string
//     }

//     var result Result

//     // Perform a JOIN query to fetch transaction details and associated item details
//     if err := s.DB.Table("transactions").
//         Select("transactions.transaction_id, transactions.client_id, transactions.client_name, transactions.item_id, items.item_name, items.description as item_description").
//         Joins("JOIN items ON transactions.item_id = items.item_id").
//         Where("transactions.transaction_id = ?", req.TransactionId).
//         First(&result).Error; err != nil {
//         if err == gorm.ErrRecordNotFound {
//             return nil, status.Errorf(codes.NotFound, "Transaction not found")
//         }
//         return nil, status.Errorf(codes.Internal, "Failed to retrieve transaction details: %v", err)
//     }

//     // Dynamically create and return the response
//     res := &pb.GetTransactionDetailsResponse{
//         TransactionId:   uint32(result.TransactionID),
//         ClientId:        result.ClientID,
//         ClientName:      result.ClientName,
//         ItemId:          uint32(result.ItemID),
//         ItemName:        result.ItemName,
//         ItemDescription: result.ItemDescription,
//     }

//     return res, nil
// }


func main() {
	config := &storage.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "root",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := models.MigrateModels(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		grpcServer := grpc.NewServer()
		pb.RegisterItemServiceServer(grpcServer, &server{DB: db})

		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port: %v", err)
		}

		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		mux := runtime.NewServeMux()
		err := pb.RegisterItemServiceHandlerFromEndpoint(
			context.Background(),
			mux,
			"localhost:50051",
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatalf("Failed to register gRPC Gateway: %v", err)
		}

		// Wrap the HTTP server with CORS middleware
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://127.0.0.1:5500", "http://localhost:5500"}, // Frontend origins
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
		}).Handler(mux)

		log.Println("Starting HTTP server on port 8080...")
		if err := http.ListenAndServe(":8080", corsHandler); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	log.Println("Received termination signal. Shutting down gracefully...")
	wg.Wait()
	log.Println("All servers stopped.")
}

