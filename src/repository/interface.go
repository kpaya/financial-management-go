package repositoryInterface

import (
	"github.com/kpaya/financial-management-go/src/domain"
)

type Transaction interface {
	InsertTransaction(input *domain.Transaction) error
}
