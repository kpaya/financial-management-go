package userRepository

import (
	"database/sql"

	"github.com/kpaya/financial-management-go/src/domain"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) InsertUser(input *domain.User) error {

	_, err := r.Db.Exec("INSERT INTO user (id, email, password, active) VALUES ($1, $2, $3, $4)", input.UserId.String(), input.Email, input.Password, input.Active)

	if err != nil {
		return err
	}

	return nil
}
