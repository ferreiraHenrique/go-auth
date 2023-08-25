package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Manager struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
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
