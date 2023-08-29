package di

import (
	"github.com/ferreiraHenrique/go-auth/adapter/http/permissionservice"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite/permissionrepository"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/domain/usecase/permissionusecase"
)

func ConfigPermissionDI(conn sqlite.PoolInterface) domain.PermissionService {
	permissionRepository := permissionrepository.New(conn)
	permissionUseCase := permissionusecase.New(permissionRepository)
	permissionService := permissionservice.New(permissionUseCase)

	return permissionService
}
