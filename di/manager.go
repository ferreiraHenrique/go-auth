package di

import (
	"github.com/ferreiraHenrique/go-auth/adapter/http/managerservice"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite/managerrepository"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/domain/usecase/managerusecase"
)

func ConfigManagerDI(conn sqlite.PoolInterface) domain.ManagerService {
	managerRepository := managerrepository.New(conn)
	managerUseCase := managerusecase.New(managerRepository)
	managerService := managerservice.New(managerUseCase)

	return managerService
}
