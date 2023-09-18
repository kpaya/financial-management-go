package transactionDto

import (
	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
)

type InputCreateNewTransactionInWallet struct {
	WalletId uuid.UUID              `json:"wallet_id"`
	Amount   float64                `json:"amount"`
	Type     domain.TransactionType `json:"type"`
}

type OutputCreateNewTransactionInWallet struct {
	Transaction domain.Transaction `json:"transaction"`
}
