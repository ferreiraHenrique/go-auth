package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Client struct {
	ID      uint
	UUID    string
	Name    string
	User    *User
	Manager *Manager
	dbRef   interface{}
}

func NewClient(id uint, uuid string, name string, user *User, manager *Manager) *Client {
	client := &Client{
		ID:      id,
		UUID:    uuid,
		Name:    name,
		User:    user,
		Manager: manager,
	}

	return client
}

func (client *Client) SetRef(ref interface{}) {
	client.dbRef = ref
}

func (client *Client) GetRef() interface{} {
	return client.dbRef
}

type ClientService interface {
	Create(response http.ResponseWriter, request *http.Request)
	AttachPermission(response http.ResponseWriter, request *http.Request)
	ListByManager(response http.ResponseWriter, request *http.Request)
}

type ClientUseCase interface {
	Create(clientRequest *dto.CreateClientRequest, managerID uint) (*Client, error)
	AttachPermission(clientRequest *dto.AttachClientPermissionRequest) error
	ListByManager(managerID uint) (*[]Client, error)
}

type ClientRepository interface {
	Create(clientRequest *dto.CreateClientRequest, managerID uint) (*Client, error)
	FindByUUID(uuid string) (*Client, error)
	AttachPermission(client *Client, permission *Permission) error
	ListByManager(managerID uint) (*[]Client, error)
}
