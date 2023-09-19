package interfaces

import transactionDto "github.com/kpaya/financial-management-go/src/dto/transaction"

type TransactionRepository interface {
	CreateNewTransactionInWallet(input *transactionDto.InputCreateNewTransactionInWallet) (*transactionDto.OutputCreateNewTransactionInWallet, error)
}
