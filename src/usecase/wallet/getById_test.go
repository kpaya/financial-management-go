package walletUseCase_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/database"
	walletDto "github.com/kpaya/financial-management-go/src/dto/wallet"
	walletRepository "github.com/kpaya/financial-management-go/src/repository/wallet"
	walletUseCase "github.com/kpaya/financial-management-go/src/usecase/wallet"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetWalletById(t *testing.T) {
	db := database.NewDbInstance()

	walletId := uuid.New()

	t.Run("create wallet", func(t *testing.T) {
		_, err := db.Exec("INSERT INTO wallet (id, user_id, active) VALUES (?, ?, ?)", walletId.String(), uuid.New().String(), true)
		if err != nil {
			t.Errorf("Error to create wallet register at database: %q\n", err)
		}
	})

	repository := walletRepository.NewWalletRepository(db)
	useCase := walletUseCase.NewGetById(repository)

	wallet, err := useCase.Execute(&walletDto.InputGetById{
		WalletId: walletId,
	})

	if err != nil {
		t.Errorf("Error to get wallet by id: %q\n", err)
	}

	switch {

	case wallet.WalletId != walletId:
		t.Errorf("Expected wallet id %q, got %q", walletId.String(), wallet.WalletId)
		break
	case wallet.UserId.String() == "":
		t.Errorf("Expected user id not empty, got %q", wallet.UserId)
		break
	case wallet.Active != true:
		t.Errorf("Expected active true, got %t", wallet.Active)
		break
	}

	t.Logf("Wallet: %v", wallet)

}
