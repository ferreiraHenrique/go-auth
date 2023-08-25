package userusecase

import (
	"errors"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (usecase usecase) Signin(userRequest *dto.SigninUserRequest) (*domain.User, error) {
	user, err := usecase.repository.FindByUsername(userRequest)
	if err != nil {
		return nil, err
	}

	if !user.Password.Compare(userRequest.Password) {
		return nil, errors.New("password doesn't match")
	}

	return user, nil
}
