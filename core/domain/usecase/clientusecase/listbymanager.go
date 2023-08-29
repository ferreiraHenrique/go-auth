package clientusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

func (usecase usecase) ListByManager(managerID uint) (*[]domain.Client, error) {
	return usecase.clientRepository.ListByManager(managerID)
}
