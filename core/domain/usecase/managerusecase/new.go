package managerusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

type usecase struct {
	repository domain.ManagerRepository
}

func New(repository domain.ManagerRepository) domain.ManagerUseCase {
	return &usecase{
		repository: repository,
	}
}
