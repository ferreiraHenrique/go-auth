package userrepository

import "github.com/ferreiraHenrique/go-auth/adapter/sqlite"

func (repository repository) IsAdmin(userID uint) bool {
	model := sqlite.Admin{UserID: userID}
	if err := repository.db.Find(&model, "user_id = ?", model.UserID).Error; err != nil {
		return false
	}

	return model.ID != 0
}
