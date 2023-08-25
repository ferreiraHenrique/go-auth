package managerusecase

import (
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (usecase usecase) Create(managerRequest *dto.CreateManagerRequest) (*domain.Manager, error) {
	manager, err := usecase.repository.Create(managerRequest)
	if err != nil {
		return nil, err
	}

	return manager, nil
}
