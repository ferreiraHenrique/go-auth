package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type CreateManagerRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromJSONCreateManagerRequest(body io.Reader) (*CreateManagerRequest, error) {
	createManagerRequest := CreateManagerRequest{}
	if err := json.NewDecoder(body).Decode(&createManagerRequest); err != nil {
		return nil, err
	}

	if createManagerRequest.Name == "" {
		return nil, errors.New("'name' is required")
	}

	if createManagerRequest.Username == "" {
		return nil, errors.New("'username' is required")
	}

	if createManagerRequest.Password == "" {
		return nil, errors.New("'password' is required")
	}

	return &createManagerRequest, nil
}
