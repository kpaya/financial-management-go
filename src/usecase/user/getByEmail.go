package userUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type GetByEmail struct {
	Repository repositoryInterface.User
}

func NewGetByEmail(repository repositoryInterface.User) (*GetByEmail, error) {
	return &GetByEmail{
		Repository: repository,
	}, nil
}

func (g *GetByEmail) Execute(input *userDto.InputGetByEmail) (*domain.User, error) {
	user, err := g.Repository.GetByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
