package domain

import "github.com/google/uuid"

type Wallet struct {
	WalletId     uuid.UUID
	UserId       uuid.UUID
	Transactions []Transaction
	Active       bool
}

func NewWallet(userId uuid.UUID) (*Wallet, error) {
	wallet := new(Wallet)
	walletId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	wallet.WalletId = walletId
	wallet.UserId = userId
	wallet.Active = true

	return wallet, nil
}
