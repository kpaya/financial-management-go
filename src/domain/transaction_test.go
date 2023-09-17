package domain_test

import (
	"testing"

	"github.com/kpaya/financial-management-go/src/domain"
)

func TestTransaction(t *testing.T) {
	transaction, err := domain.NewTransaction(300.23, domain.Deposit)

	if err != nil {
		t.Errorf("there was an error creating the transaction: %v", err)
	}

	if transaction.Value != 300.23 {
		t.Errorf("expected value to be 300.23, got %v", transaction.Value)
	}

	if transaction.Type != domain.Deposit {
		t.Errorf("expected type to be Deposit, got %v", transaction.Type)
	}
}
