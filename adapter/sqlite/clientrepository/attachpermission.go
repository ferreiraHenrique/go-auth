package clientrepository

import (
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/core/domain"
)

func (repository repository) AttachPermission(client *domain.Client, permission *domain.Permission) error {
	clientModel := client.GetRef().(*sqlite.Client)
	permissionModel := permission.GetRef().(*sqlite.Permission)

	return repository.db.Model(clientModel).Association("Permissions").Append(permissionModel)
}
