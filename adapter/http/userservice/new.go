package userservice

import "github.com/ferreiraHenrique/go-auth/core/domain"

type service struct {
	usercase domain.UserUseCase
}

func New(usercase domain.UserUseCase) domain.UserService {
	return &service{
		usercase: usercase,
	}
}
