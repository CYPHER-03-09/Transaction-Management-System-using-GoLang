package models

import "gorm.io/gorm"

// type Item struct {
// 	ItemID      uint   `gorm:"primaryKey;autoIncrement" json:"item_id"`
// 	ItemName    string `gorm:"not null" json:"item_name"`
// 	Description string `gorm:"not null" json:"description"`
// 	Transactions []Transaction `gorm:"foreignKey:ItemID"` // Relationship with Transactions table
// }


// type Transaction struct {
// 	TransactionID uint   `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
// 	ClientID      string `gorm:"not null" json:"client_id"`
// 	ClientName    string `gorm:"not null" json:"client_name"`
// 	ItemID        uint   `gorm:"not null" json:"item_id"` // Foreign key from Items table
// }

// func MigrateModels(db *gorm.DB) error {
// 	// Migrate both tables together
// 	return db.AutoMigrate(&Item{}, &Transaction{})
// }


// Item model
type Item struct {
	ItemID      uint   `gorm:"primaryKey;autoIncrement" json:"item_id"`
	ItemName    string `gorm:"not null" json:"item_name"`
	Description string `gorm:"not null" json:"description"`
	Transactions []Transaction `gorm:"foreignKey:ItemID"` // Relationship with Transactions table
}

// Client model
type Client struct {
	ClientID   uint   `gorm:"primaryKey;autoIncrement" json:"client_id"`
	ClientName string `gorm:"not null" json:"client_name"`
	Phone      string `gorm:"not null" json:"phone"`
	Email      string `gorm:"not null" json:"email"`
	Transactions []Transaction `gorm:"foreignKey:ClientID"` // Relationship with Transactions table
}

// Transaction model
type Transaction struct {
	TransactionID uint   `gorm:"primaryKey;autoIncrement" json:"transaction_id"`
	ClientID      uint   `gorm:"not null" json:"client_id"`
	Client        Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Foreign key reference to Client
	ItemID        uint   `gorm:"not null" json:"item_id"`
	Item          Item   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`   // Foreign key reference to Item
}


func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&Item{}, &Client{}, &Transaction{})
}

