package permissionusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

func (usecase usecase) ListByManager(managerID uint) (*[]domain.Permission, error) {
	permissions, err := usecase.repository.ListByManager(managerID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
