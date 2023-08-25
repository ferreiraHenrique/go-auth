package managerrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
	"gorm.io/gorm"
)

func (repository repository) Create(managerRequest *dto.CreateManagerRequest) (*domain.Manager, error) {
	password := domain.NewPassword(managerRequest.Password)
	password.Hash()
	user := sqlite.User{
		Username: managerRequest.Username,
		Password: password.Password,
	}
	manager := sqlite.Manager{Name: managerRequest.Name}

	if err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		manager.UserID = user.ID
		if err := tx.Create(&manager).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return domain.NewManager(
		manager.ID,
		manager.Name,
		domain.NewUser(user.ID, user.Username, user.Password),
	), nil
}
