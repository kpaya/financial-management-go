package repositoryInterface

import (
	"github.com/kpaya/financial-management-go/src/domain"
)

type Transaction interface {
	InsertTransaction(input *domain.Transaction) error
}

type User interface {
	Insert(input *domain.User) error
	GetById(uuid string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
