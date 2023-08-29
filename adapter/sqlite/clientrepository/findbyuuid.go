package clientrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) FindByUUID(uuid string) (*domain.Client, error) {
	var model sqlite.Client
	if err := repository.db.Find(&model, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	client := domain.NewClient(
		model.ID,
		model.UUID,
		model.Name,
		nil,
		nil,
	)
	client.SetRef(&model)

	return client, nil
}
