package managerservice

import "github.com/ferreiraHenrique/go-auth/core/domain"

type service struct {
	usecase domain.ManagerUseCase
}

func New(usecase domain.ManagerUseCase) domain.ManagerService {
	return &service{
		usecase: usecase,
	}
}
