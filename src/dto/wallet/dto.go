package walletDto

import "github.com/google/uuid"

type InputNewCreateWallet struct {
	UserId string
}

type OutputNewCreateWallet struct {
	WalletId uuid.UUID
}

type InputGetById struct {
	WalletId uuid.UUID
}
