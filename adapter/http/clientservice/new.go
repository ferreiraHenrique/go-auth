package clientservice

import "github.com/ferreiraHenrique/go-auth/core/domain"

type service struct {
	usecaase domain.ClientUseCase
}

func New(usecase domain.ClientUseCase) domain.ClientService {
	return &service{
		usecaase: usecase,
	}
}
