package permissionusecase

import (
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (usecase usecase) Create(permissionRequest *dto.CreatePermissionRequest, managerID uint) (*domain.Permission, error) {
	permission, err := usecase.repository.Create(permissionRequest, managerID)
	if err != nil {
		return nil, err
	}

	return permission, nil
}
