package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Manager struct {
	ID   uint
	Name string
	User *User
}

func NewManager(id uint, name string, user *User) *Manager {
	manager := &Manager{
		ID:   id,
		Name: name,
		User: user,
	}

	return manager
}

type ManagerService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ManagerUseCase interface {
	Create(managerRequest *dto.CreateManagerRequest) (*Manager, error)
}

type ManagerRepository interface {
	Create(managerRequest *dto.CreateManagerRequest) (*Manager, error)
}
