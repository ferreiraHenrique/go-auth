package domain

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

type Client struct {
	ID   uint
	Name string
	User *User
}

func NewClient(id uint, name string, user *User) *Client {
	client := &Client{
		ID:   id,
		Name: name,
		User: user,
	}

	return client
}

type ClientService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ClientUseCase interface {
	Create(clientRequest *dto.CreateClientRequest) (*Client, error)
}

type ClientRepository interface {
	Create(clientRequest *dto.CreateClientRequest) (*Client, error)
}
