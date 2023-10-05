package walletUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	walletDto "github.com/kpaya/financial-management-go/src/dto/wallet"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type GetById struct {
	Repository repositoryInterface.Wallet
}

func NewGetById(repository repositoryInterface.Wallet) *GetById {
	return &GetById{
		Repository: repository,
	}
}

func (w *GetById) Execute(input *walletDto.InputGetById) (*domain.Wallet, error) {
	wallet, err := w.Repository.GetById(input.WalletId.String())
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
