package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type User struct {
	ID       uint
	Username string
	Password *Password
	Role     string
}

func NewUser(id uint, username, password string) *User {
	user := &User{
		ID:       id,
		Username: username,
		Password: NewPassword(password),
	}

	return user
}

type UserService interface {
	Signin(response http.ResponseWriter, request *http.Request)
	IsAdmin(h http.Handler) http.Handler
	IsManager(h http.Handler) http.Handler
}

type UserUseCase interface {
	Signin(userRequest *dto.SigninUserRequest) (*User, error)
	IsAdmin(signedToken string) bool
	IsManager(signedToken string) bool
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
	IsAdmin(userID uint) bool
	IsManager(userID uint) bool
}
