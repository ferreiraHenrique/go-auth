package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Client struct {
	ID      uint
	Name    string
	User    *User
	Manager *Manager
}

func NewClient(id uint, name string, user *User, manager *Manager) *Client {
	client := &Client{
		ID:      id,
		Name:    name,
		User:    user,
		Manager: manager,
	}

	return client
}

type ClientService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ClientUseCase interface {
	Create(clientRequest *dto.CreateClientRequest, managerID uint) (*Client, error)
}

type ClientRepository interface {
	Create(clientRequest *dto.CreateClientRequest, managerID uint) (*Client, error)
}
