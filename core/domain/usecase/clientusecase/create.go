package clientusecase

import (
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (usecase usecase) Create(clientRequest *dto.CreateClientRequest) (*domain.Client, error) {
	client, err := usecase.repository.Create(clientRequest)
	if err != nil {
		return nil, err
	}

	return client, nil
}
