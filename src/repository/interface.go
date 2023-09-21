package repositoryInterface

import (
	"github.com/kpaya/financial-management-go/src/domain"
)

type Transaction interface {
	InsertTransaction(input *domain.Transaction) error
}

type User interface {
	Insert(input *domain.User) error
	Get(uuid string) (*domain.User, error)
}
