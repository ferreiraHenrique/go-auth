package clientusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

type usecase struct {
	repository domain.ClientRepository
}

func New(repository domain.ClientRepository) domain.ClientUseCase {
	return &usecase{
		repository: repository,
	}
}
