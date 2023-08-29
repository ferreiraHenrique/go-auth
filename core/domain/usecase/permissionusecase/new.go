package permissionusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

type usecase struct {
	repository domain.PermissionUseCase
}

func New(repository domain.PermissionRepository) domain.PermissionUseCase {
	return &usecase{
		repository: repository,
	}
}
