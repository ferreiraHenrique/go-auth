package clientservice

import "github.com/ferreiraHenrique/go-auth/core/domain"

type service struct {
	usecase domain.ClientUseCase
}

func New(usecase domain.ClientUseCase) domain.ClientService {
	return &service{
		usecase: usecase,
	}
}
