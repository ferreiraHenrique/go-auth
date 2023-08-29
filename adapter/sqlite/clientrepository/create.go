package clientrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
	"gorm.io/gorm"
)

func (repository repository) Create(clientRequest *dto.CreateClientRequest, managerID uint) (*domain.Client, error) {
	password := domain.NewPassword(clientRequest.Password)
	password.Hash()
	user := sqlite.User{
		Username: clientRequest.Username,
		Password: password.Password,
	}
	client := sqlite.Client{
		Name:      clientRequest.Name,
		ManagerID: managerID,
	}

	if err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		client.UserID = user.ID
		if err := tx.Create(&client).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return domain.NewClient(
		client.ID,
		client.UUID,
		client.Name,
		domain.NewUser(user.ID, user.Username, user.Password),
		nil,
	), nil
}
