package userUseCase

import (
	"github.com/kpaya/financial-management-go/src/domain"
	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	repositoryInterface "github.com/kpaya/financial-management-go/src/repository"
)

type GetById struct {
	Repository repositoryInterface.User
}

func NewGetById(repository repositoryInterface.User) (*GetById, error) {
	return &GetById{
		Repository: repository,
	}, nil
}

func (g *GetById) Execute(input *userDto.InputGetById) (*domain.User, error) {
	user, err := g.Repository.GetById(input.UserId.String())
	if err != nil {
		return nil, err
	}

	return user, nil
}
