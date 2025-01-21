// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ItemService_AddItem_FullMethodName               = "/ItemService/AddItem"
	ItemService_GetItems_FullMethodName              = "/ItemService/GetItems"
	ItemService_AddClient_FullMethodName             = "/ItemService/AddClient"
	ItemService_GetClients_FullMethodName            = "/ItemService/GetClients"
	ItemService_AddTransaction_FullMethodName        = "/ItemService/AddTransaction"
	ItemService_GetTransactions_FullMethodName       = "/ItemService/GetTransactions"
	ItemService_GetTransactionDetails_FullMethodName = "/ItemService/GetTransactionDetails"
)

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	// Add an item
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	// Get all items
	GetItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetItemsResponse, error)
	// Add a client
	AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error)
	// Get all clients
	GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetClientsResponse, error)
	// Add a transaction
	AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error)
	// Get all transactions
	GetTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	// Get transaction details by ID
	GetTransactionDetails(ctx context.Context, in *GetTransactionDetailsRequest, opts ...grpc.CallOption) (*GetTransactionDetailsResponse, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, ItemService_AddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, ItemService_GetItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddClientResponse)
	err := c.cc.Invoke(ctx, ItemService_AddClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetClientsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClientsResponse)
	err := c.cc.Invoke(ctx, ItemService_GetClients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTransactionResponse)
	err := c.cc.Invoke(ctx, ItemService_AddTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, ItemService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetTransactionDetails(ctx context.Context, in *GetTransactionDetailsRequest, opts ...grpc.CallOption) (*GetTransactionDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionDetailsResponse)
	err := c.cc.Invoke(ctx, ItemService_GetTransactionDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility.
type ItemServiceServer interface {
	// Add an item
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	// Get all items
	GetItems(context.Context, *Empty) (*GetItemsResponse, error)
	// Add a client
	AddClient(context.Context, *AddClientRequest) (*AddClientResponse, error)
	// Get all clients
	GetClients(context.Context, *Empty) (*GetClientsResponse, error)
	// Add a transaction
	AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error)
	// Get all transactions
	GetTransactions(context.Context, *Empty) (*GetTransactionsResponse, error)
	// Get transaction details by ID
	GetTransactionDetails(context.Context, *GetTransactionDetailsRequest) (*GetTransactionDetailsResponse, error)
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedItemServiceServer struct{}

func (UnimplementedItemServiceServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedItemServiceServer) GetItems(context.Context, *Empty) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedItemServiceServer) AddClient(context.Context, *AddClientRequest) (*AddClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (UnimplementedItemServiceServer) GetClients(context.Context, *Empty) (*GetClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedItemServiceServer) AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedItemServiceServer) GetTransactions(context.Context, *Empty) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedItemServiceServer) GetTransactionDetails(context.Context, *GetTransactionDetailsRequest) (*GetTransactionDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionDetails not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}
func (UnimplementedItemServiceServer) testEmbeddedByValue()                     {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	// If the following call pancis, it indicates UnimplementedItemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_AddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItems(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_AddClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).AddClient(ctx, req.(*AddClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetClients(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_AddTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).AddTransaction(ctx, req.(*AddTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetTransactions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetTransactionDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetTransactionDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetTransactionDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetTransactionDetails(ctx, req.(*GetTransactionDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ItemService_AddItem_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _ItemService_GetItems_Handler,
		},
		{
			MethodName: "AddClient",
			Handler:    _ItemService_AddClient_Handler,
		},
		{
			MethodName: "GetClients",
			Handler:    _ItemService_GetClients_Handler,
		},
		{
			MethodName: "AddTransaction",
			Handler:    _ItemService_AddTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _ItemService_GetTransactions_Handler,
		},
		{
			MethodName: "GetTransactionDetails",
			Handler:    _ItemService_GetTransactionDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
