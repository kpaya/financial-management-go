package transactionRepository

import (
	"database/sql"

	transactionDto "github.com/kpaya/financial-management-go/src/dto/transaction"
)

type TransactionRepository struct {
	Db *sql.DB
}

func NewTransactionRepository(db *sql.DB) (*TransactionRepository, error) {
	return &TransactionRepository{
		Db: db,
	}, nil
}

func (t *TransactionRepository) CreateNewTransactionInWallet(input *transactionDto.InputCreateNewTransactionInWallet) (*transactionDto.OutputCreateNewTransactionInWallet, error) {
	return nil, nil
}
