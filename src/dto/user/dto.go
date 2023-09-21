package userDto

import "github.com/google/uuid"

type InputNewCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OutputNewCreateUser struct {
	UserId uuid.UUID `json:"user_id"`
}

type InputGetUserById struct {
	UserId uuid.UUID `json:"user_id"`
}
