package userrepository

import (
	"errors"

	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) FindByUsername(username string) (*domain.User, error) {
	model := sqlite.User{Username: username}
	if err := repository.db.Find(&model, "username = ?", model.Username).Error; err != nil {
		return nil, err
	}

	user := domain.NewUser(model.ID, model.Username, model.Password)
	if user == nil {
		return nil, errors.New("user not found")
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}
