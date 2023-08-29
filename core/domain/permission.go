package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Permission struct {
	ID      uint
	UUID    string
	Name    string
	Manager *Manager
}

func NewPermission(id uint, uuid string, name string, manager *Manager) *Permission {
	return &Permission{
		ID:      id,
		UUID:    uuid,
		Name:    name,
		Manager: manager,
	}
}

type PermissionService interface {
	Create(response http.ResponseWriter, request *http.Request)
	ListByManager(response http.ResponseWriter, request *http.Request)
}

type PermissionUseCase interface {
	Create(permissionRequest *dto.CreatePermissionRequest, managerID uint) (*Permission, error)
	ListByManager(managerID uint) (*[]Permission, error)
}

type PermissionRepository interface {
	Create(permissionRequest *dto.CreatePermissionRequest, managerID uint) (*Permission, error)
	ListByManager(managerID uint) (*[]Permission, error)
}
