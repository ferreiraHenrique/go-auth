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

type ClientService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ClientUseCase interface {
	Create(clientRequest *dto.CreateClientRequest) (*Client, error)
}

type ClientRepository interface {
	Create(clientRequest *dto.CreateClientRequest) (*Client, error)
}
