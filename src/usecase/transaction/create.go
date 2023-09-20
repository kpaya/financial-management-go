package transactionUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	transactionDto "github.com/kpaya/financial-management-go/src/dto/transaction"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type CreateNewTransaction struct {
	TransactionRepository repositoryInterface.Transaction
}

func NewCreateNewTransaction(repository repositoryInterface.Transaction) (*CreateNewTransaction, error) {
	return &CreateNewTransaction{TransactionRepository: repository}, nil
}

func (useCase *CreateNewTransaction) Execute(input *transactionDto.InputInsertTransaction) (*transactionDto.OutputInsertTransaction, error) {
	transaction, err := domain.NewTransaction(input.Amount, input.Type)

	if err != nil {
		return nil, err
	}

	err = useCase.TransactionRepository.InsertTransaction(transaction)
	if err != nil {
		return nil, err
	}

	output := &transactionDto.OutputInsertTransaction{
		WalletId: transaction.TransactionId,
	}

	return output, nil
}
