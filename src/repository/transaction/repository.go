package transactionRepository

import (
	"database/sql"

	"github.com/kpaya/financial-management-go/src/domain"
)

type TransactionRepository struct {
	Db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		Db: db,
	}
}

func (t *TransactionRepository) InsertTransaction(input *domain.Transaction) error {
	_, err := t.Db.Exec("INSERT INTO transactions (id, balance, type) VALUES (?, ?, ?)", input.TransactionId.String(), input.Value, input.Type)

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepository) GetTransactionsByWalletId(walletId string) (*domain.Transaction, error) {
	row := t.Db.QueryRow("SELECT id, balance, type FROM transactions WHERE id = ?", walletId)

	var transaction *domain.Transaction
	err := row.Scan(&transaction.TransactionId, &transaction.Value, &transaction.Type)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
