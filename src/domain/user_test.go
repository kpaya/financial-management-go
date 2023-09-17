package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kpaya/financial-management-go/src/domain"
	"golang.org/x/crypto/bcrypt"
)

func TestUser(t *testing.T) {
	email, password := "test@test.com", "123456"

	user, err := domain.NewUser(email, password)
	if err != nil {
		t.Errorf("Error creating user: %s", err)
	}

	if user.UserId == uuid.Nil {
		t.Error("User ID is nil")
	}

	if len(user.Email) <= 0 {
		t.Error("User email is empty")
	}

	if user.Email != email {
		t.Errorf("User email is not %s", email)
	}

	if bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); bcryptErr != nil {
		t.Errorf(`The user hash password: "%s" doesn't fit with the input password: "%s"`, user.Password, password)
	}
}
