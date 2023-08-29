package permissionrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) FindByUUID(uuid string) (*domain.Permission, error) {
	var model sqlite.Permission
	if err := repository.db.Find(&model, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	permission := domain.NewPermission(
		model.ID,
		model.UUID,
		model.Name,
		nil,
	)
	permission.SetRef(&model)

	return permission, nil
}
