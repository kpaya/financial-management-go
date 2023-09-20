package transactionDto

import (
	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
)

type InputInsertTransaction struct {
	WalletId uuid.UUID              `json:"wallet_id"`
	Amount   float64                `json:"amount"`
	Type     domain.TransactionType `json:"type"`
}

type OutputInsertTransaction struct {
	WalletId uuid.UUID `json:"wallet_id"`
}
