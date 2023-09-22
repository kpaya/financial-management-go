package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   uuid.UUID
	Email    string
	Password string
	Active   bool
}

func NewUser(email, password string) (*User, error) {
	user := new(User)

	userID := uuid.New()

	user.UserId = userID
	user.Email = email
	user.Active = true
	passwordHash, err := EncryptPassword(password)

	if err != nil {
		return nil, err
	}

	user.Password = passwordHash

	return user, nil
}

func EncryptPassword(password string) (string, error) {
	if len(password) <= 0 || len(password) > 72 {
		return "", errors.New("the password must be between 1 and 72 characters")
	}

	password = strings.TrimSpace(password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
