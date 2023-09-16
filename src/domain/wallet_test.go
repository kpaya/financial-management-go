package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
)

func TestWallet(t *testing.T) {
	userId, err := uuid.NewUUID()
	if err != nil {
		t.Error("Error generating UUID")
	}

	wallet, err := domain.NewWallet(userId)

	if err != nil {
		t.Error("Error creating wallet")
	}

	if wallet.WalletId == uuid.Nil {
		t.Error("Wallet ID is nil")
	}

	if wallet.UserId == uuid.Nil {
		t.Error("User ID is nil")
	}

	if wallet.Active != true {
		t.Error("Wallet is not active")
	}
}
