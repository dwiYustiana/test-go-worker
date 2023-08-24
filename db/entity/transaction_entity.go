package entity

import (
	"time"

	"github.com/google/uuid"
)

// struct
type Transaction struct {
	ID            uuid.UUID `gorm:"primary_key" json:"id"'`
	TransactionId int       `json:"transaction_id"`
	RequestId     int       `json:"request_id"`
	Customer      string    `json:"customer"`
	Price         float32   `json:"price"`
	Quantity      int       `json:"quantity"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e *Transaction) TableName() string {
	return "transactions"
}
