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

func (r *UserRepository) Insert(input *domain.User) error {

	_, err := r.Db.Exec("INSERT INTO user (id, email, password, active) VALUES ($1, $2, $3, $4)", input.UserId.String(), input.Email, input.Password, input.Active)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(uuid string) (*domain.User, error) {
	var user domain.User

	err := r.Db.QueryRow("SELECT id, email, password, active FROM user WHERE id = $1", uuid).Scan(&user.UserId, &user.Email, &user.Password, &user.Active)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.Db.QueryRow("SELECT id, email, password, active FROM user WHERE email = $1", email).Scan(&user.UserId, &user.Email, &user.Password, &user.Active)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
