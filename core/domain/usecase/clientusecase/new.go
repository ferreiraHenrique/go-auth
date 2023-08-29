package clientusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

type usecase struct {
	clientRepository     domain.ClientRepository
	permissionRepository domain.PermissionRepository
}

func New(
	clientRepository domain.ClientRepository,
	permissionRepository domain.PermissionRepository,
) domain.ClientUseCase {
	return &usecase{
		clientRepository:     clientRepository,
		permissionRepository: permissionRepository,
	}
}
