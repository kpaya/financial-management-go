package userUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type GetUserById struct {
	Repository repositoryInterface.User
}

func NewGetUserById(repository repositoryInterface.User) (*GetUserById, error) {
	return &GetUserById{
		Repository: repository,
	}, nil
}

func (g *GetUserById) Execute(input *userDto.InputGetUserById) (*domain.User, error) {
	user, err := g.Repository.Get(input.UserId.String())
	if err != nil {
		return nil, err
	}

	return user, nil
}
