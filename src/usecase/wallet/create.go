package walletUseCase

import (
	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
	walletDto "github.com/kpaya/financial-management-go/src/dto/wallet"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type CreateNewWallet struct {
	Repository repositoryInterface.Wallet
}

func NewCreateNewWallet(repository repositoryInterface.Wallet) (*CreateNewWallet, error) {
	return &CreateNewWallet{Repository: repository}, nil
}

func (u *CreateNewWallet) Execute(input *walletDto.InputNewCreateWallet) (*walletDto.OutputNewCreateWallet, error) {
	userId, err := uuid.Parse(input.UserId)
	if err != nil {
		return nil, err
	}
	wallet, err := domain.NewWallet(uuid.Nil, userId, []domain.Transaction{}, true)

	err = u.Repository.Insert(wallet)

	if err != nil {
		return nil, err
	}
	return &walletDto.OutputNewCreateWallet{WalletId: wallet.WalletId}, nil
}
