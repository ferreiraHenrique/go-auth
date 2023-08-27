package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type CreateClientRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromJSONCreateClientRequest(body io.Reader) (*CreateClientRequest, error) {
	createClientRequest := CreateClientRequest{}
	if err := json.NewDecoder(body).Decode(&createClientRequest); err != nil {
		return nil, err
	}

	if createClientRequest.Name == "" {
		return nil, errors.New("'name' is required")
	}

	if createClientRequest.Username == "" {
		return nil, errors.New("'username' is required")
	}

	if createClientRequest.Password == "" {
		return nil, errors.New("'password' is required")
	}

	return &createClientRequest, nil
}
