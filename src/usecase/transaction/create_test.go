package transactionUseCase_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
	transactionDto "github.com/kpaya/financial-management-go/src/dto/transaction"
	transactionRepository "github.com/kpaya/financial-management-go/src/repository/transaction"
	transactionUseCase "github.com/kpaya/financial-management-go/src/usecase/transaction"
	_ "github.com/mattn/go-sqlite3"
)

func TestCreateNewTransactionInWallet(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../database_test.db")
	if err != nil {
		t.Errorf("Error opening database: %q\n", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, balance REAL, type TEXT)")
	if err != nil {
		t.Errorf("Error to exec query %q\n", err)
	}

	repository := transactionRepository.NewTransactionRepository(db)
	if err != nil {
		t.Error(err)
	}

	useCase, err := transactionUseCase.NewCreateNewTransaction(repository)
	if err != nil {
		t.Error(err)
	}

	output, err := useCase.Execute(&transactionDto.InputInsertTransaction{
		WalletId: uuid.New(),
		Amount:   101.53,
		Type:     domain.Deposit,
	})

	if err != nil {
		t.Error(err)
	}

	log.Println(output)

}
