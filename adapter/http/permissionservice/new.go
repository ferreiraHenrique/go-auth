package permissionservice

import "github.com/ferreiraHenrique/go-auth/core/domain"

type service struct {
	usecase domain.PermissionUseCase
}

func New(usecase domain.PermissionUseCase) domain.PermissionService {
	return &service{
		usecase: usecase,
	}
}
