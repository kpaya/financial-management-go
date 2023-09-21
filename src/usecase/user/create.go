package userUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type CreateNewUser struct {
	Repository repositoryInterface.User
}

func NewCreateNewUser(repository repositoryInterface.User) (*CreateNewUser, error) {
	return &CreateNewUser{
		Repository: repository,
	}, nil
}

func (u *CreateNewUser) Execute(input *userDto.InputNewCreateUser) (*userDto.OutputNewCreateUser, error) {
	user, err := domain.NewUser(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	err = u.Repository.Insert(user)
	if err != nil {
		return nil, err
	}

	output := &userDto.OutputNewCreateUser{
		UserId: user.UserId,
	}

	return output, nil
}
