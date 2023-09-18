package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
)

func TestWalletDomain(t *testing.T) {
	walletId := uuid.New()

	userId := uuid.New()

	wallet, err := domain.NewWallet(walletId, userId, []domain.Transaction{}, true)

	if err != nil {
		t.Error("Error creating wallet")
	}

	if _, err := uuid.Parse(wallet.WalletId.String()); err != nil {
		t.Error("Wallet ID is nil")
	}

	if _, err := uuid.Parse(wallet.UserId.String()); err != nil {
		t.Error("User ID is nil")
	}

	if wallet.Active != true {
		t.Error("Wallet is not active")
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
