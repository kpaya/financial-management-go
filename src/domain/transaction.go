package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	Withdraw TransactionType = "W"
	Deposit  TransactionType = "D"
)

type Transaction struct {
	TransactionId uuid.UUID
	Value         float64
	Type          TransactionType
	CreatedAt     time.Time
}

func NewTransaction(value float64, transactionType TransactionType) (*Transaction, error) {
	transaction := new(Transaction)
	transactionId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	transaction.TransactionId = transactionId
	transaction.Value = value
	transaction.Type = transactionType
	transaction.CreatedAt = time.Now()

	return transaction, nil
}
