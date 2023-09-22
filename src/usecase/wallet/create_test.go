package walletUseCase_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	walletDto "github.com/kpaya/financial-management-go/src/dto/wallet"
	walletRepository "github.com/kpaya/financial-management-go/src/repository/wallet"
	walletUseCase "github.com/kpaya/financial-management-go/src/usecase/wallet"
	_ "github.com/mattn/go-sqlite3"
)

func TestInsertWallet(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../database_test.db")
	if err != nil {
		t.Errorf("Error opening database: %q\n", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS wallet (id TEXT PRIMARY KEY, user_id TEXT, active BOOLEAN, FOREIGN KEY(user_id) REFERENCES user(id))")
	if err != nil {
		t.Errorf("Error to exec query %q\n", err)
	}

	userId := uuid.New()

	repository := walletRepository.NewWalletRepository(db)

	useCase, err := walletUseCase.NewCreateNewWallet(repository)

	if err != nil {
		t.Errorf("Error to create usecase %q\n", err)
	}

	input := new(walletDto.InputNewCreateWallet)
	input.UserId = userId.String()

	output, err := useCase.Execute(input)

	if err != nil {
		t.Errorf("Error to execute usecase %q\n", err)
	}

	if _, err := uuid.Parse(output.WalletId.String()); err != nil {
		t.Errorf("Error to parse uuid %q\n", err)
	}

	log.Default().Println(output)
}
