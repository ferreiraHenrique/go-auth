package managerrepository

import (
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (repository repository) Create(managerRequest *dto.CreateManagerRequest) (*domain.Manager, error) {
	manager := domain.Manager{Name: managerRequest.Name}

	if err := repository.db.Create(&manager).Error; err != nil {
		return nil, err
	}

	return &manager, nil
}
