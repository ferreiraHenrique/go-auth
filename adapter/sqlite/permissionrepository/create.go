package permissionrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (repository repository) Create(permissionRequest *dto.CreatePermissionRequest, managerID uint) (*domain.Permission, error) {
	permission := sqlite.Permission{
		Name:      permissionRequest.Name,
		ManagerID: managerID,
	}

	if err := repository.db.Create(&permission).Error; err != nil {
		return nil, err
	}

	return domain.NewPermission(
		permission.ID,
		permission.Name,
		nil,
	), nil
}
