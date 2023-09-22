package userUseCase_test

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
	userDto "github.com/kpaya/financial-management-go/src/dto/user"
	userRepository "github.com/kpaya/financial-management-go/src/repository/user"
	userUseCase "github.com/kpaya/financial-management-go/src/usecase/user"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func TestGetById(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../database_test.db")
	if err != nil {
		t.Errorf("Error opening database: %q\n", err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id TEXT PRIMARY KEY, email TEXT, password TEXT, active BOOLEAN)")
	if err != nil {
		t.Errorf("Error to exec query %q\n", err)
	}

	email := fmt.Sprintf("%s@gmail.com", strconv.FormatInt(rand.Int63n(100000000), 10))
	password := strconv.FormatInt(rand.Int63n(100000000), 10)

	repository := userRepository.NewUserRepository(db)

	var userId uuid.UUID
	t.Run("Creating a New User", func(t *testing.T) {

		user := &userDto.InputNewCreateUser{
			Email:    email,
			Password: password,
		}

		if err != nil {
			t.Errorf("Error to create user %q\n", err)
		}

		usecase, err := userUseCase.NewCreateNewUser(repository)

		if err != nil {
			t.Errorf("Error to create usecase %q\n", err)
		}

		output, err := usecase.Execute(user)

		if err != nil {
			t.Errorf("Error to execute usecase %q\n", err)
		}

		log.Printf("Output of InsertUser: %q\n", output)

		userId = output.UserId
	})

	usecase, err := userUseCase.NewGetById(repository)

	if err != nil {
		t.Errorf("Error to create usecase %q\n", err)
	}

	user, err := usecase.Execute(&userDto.InputGetById{UserId: userId})

	switch {
	case err != nil:
		t.Errorf("Error to get user %q\n", err)

	case user.UserId != userId:
		t.Errorf("UserId is not the same: %q\n", err)

	case user.Email == "":
		t.Errorf("Email is not the same: %q\n", err)

	case user.Password == "":
		t.Errorf("Password is null: %q\n", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		t.Errorf("Error to compare password %q\n", err)
	}

	log.Default().Println("User: ", user)
}
