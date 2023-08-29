package clientrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) ListByManager(managerID uint) (*[]domain.Client, error) {
	var cs []sqlite.Client
	if err := repository.db.Find(&cs, "manager_id = ?", managerID).Error; err != nil {
		return nil, err
	}

	var clients []domain.Client
	for _, c := range cs {
		clients = append(clients, *domain.NewClient(
			c.ID,
			c.UUID,
			c.Name,
			nil,
			nil,
		))
	}

	return &clients, nil
}
