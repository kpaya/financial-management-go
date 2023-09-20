package userUseCase_test

import (
	"database/sql"
	"log"
	"testing"

	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	userRepository "github.com/kpaya/financial-management-go/src/repository/user"
	userUseCase "github.com/kpaya/financial-management-go/src/usecase/user"
	_ "github.com/mattn/go-sqlite3"
)

func TestInsertUser(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../database_test.db")
	if err != nil {
		t.Errorf("Error opening database: %q\n", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id TEXT PRIMARY KEY, email TEXT, password TEXT, active BOOLEAN)")
	if err != nil {
		t.Errorf("Error to exec query %q\n", err)
	}

	email := "test@gmail.com"
	password := "12345678"

	user := &userDto.InputNewCreateUser{
		Email:    email,
		Password: password,
	}

	if err != nil {
		t.Errorf("Error to create user %q\n", err)
	}

	repository := userRepository.NewUserRepository(db)

	usecase, err := userUseCase.NewCreateNewUser(repository)

	if err != nil {
		t.Errorf("Error to create usecase %q\n", err)
	}

	output, err := usecase.Execute(user)

	if err != nil {
		t.Errorf("Error to execute usecase %q\n", err)
	}

	log.Printf("Output of InsertUser: %q\n", output)
}
