package clientusecase

import (
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (usecase usecase) Create(clientRequest *dto.CreateClientRequest, managerID uint) (*domain.Client, error) {
	client, err := usecase.clientRepository.Create(clientRequest, managerID)
	if err != nil {
		return nil, err
	}

	return client, nil
}
