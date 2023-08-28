package userrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) IsManager(userID uint) (bool, *domain.Manager) {
	model := sqlite.Manager{UserID: userID}
	if err := repository.db.Find(&model, "user_id = ?", model.UserID).Error; err != nil {
		return false, nil
	}

	if model.ID == 0 {
		return false, nil
	}

	return true, domain.NewManager(
		model.ID,
		model.Name,
		domain.NewUser(
			userID,
			model.User.Username,
			model.User.Password,
		),
	)
}
