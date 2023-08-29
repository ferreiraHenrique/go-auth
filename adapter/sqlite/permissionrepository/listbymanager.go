package permissionrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) ListByManager(managerID uint) (*[]domain.Permission, error) {
	var ps []sqlite.Permission
	if err := repository.db.Find(&ps, "manager_id = ?", managerID).Error; err != nil {
		return nil, err
	}

	var permissions []domain.Permission
	for _, p := range ps {
		permissions = append(permissions, *domain.NewPermission(
			p.ID,
			p.UUID,
			p.Name,
			nil,
		))
	}

	return &permissions, nil
}
