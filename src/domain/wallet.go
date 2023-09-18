package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Wallet struct {
	WalletId     uuid.UUID
	UserId       uuid.UUID
	Transactions []Transaction
	Active       bool
}

func NewWallet(walletId, userId uuid.UUID, transactions []Transaction, active bool) (*Wallet, error) {
	wallet := new(Wallet)
	if _, err := uuid.Parse(walletId.String()); err != nil {
		wallet.WalletId = uuid.New()
	}

	if _, err := uuid.Parse(userId.String()); err != nil {
		return nil, errors.New("User ID is nil")
	}

	wallet.UserId = userId
	wallet.Transactions = transactions
	wallet.Active = active

	return wallet, nil
}
