package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type User struct {
	ID       int32
	Username string
	Password *Password
}

func NewUser(id int32, username, password string) *User {
	user := &User{
		ID:       id,
		Username: username,
		Password: NewPassword(password),
	}

	if user.Password == nil {
		return nil
	}

	return user
}

type UserService interface {
	Signin(response http.ResponseWriter, request *http.Request)
}

type UserUseCase interface {
	Signin(userRequest *dto.SigninUserRequest) (*User, error)
}

type UserRepository interface {
	FindByUsername(userRequest *dto.SigninUserRequest) (*User, error)
}
