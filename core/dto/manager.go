package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type CreateManagerRequest struct {
	Name string `json:"name"`
}

func FromJSONCreateManagerRequest(body io.Reader) (*CreateManagerRequest, error) {
	createManagerRequest := CreateManagerRequest{}
	if err := json.NewDecoder(body).Decode(&createManagerRequest); err != nil {
		return nil, err
	}

	if createManagerRequest.Name == "" {
		return nil, errors.New("'name' is required")
	}

	return &createManagerRequest, nil
}
