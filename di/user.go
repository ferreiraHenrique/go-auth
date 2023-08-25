package di

import (
	"github.com/ferreiraHenrique/go-auth/adapter/http/userservice"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite/userrepository"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/domain/usecase/userusecase"
)

func ConfigUserDI(conn sqlite.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)

	return userService
}
