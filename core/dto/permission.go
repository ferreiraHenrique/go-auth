package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type CreatePermissionRequest struct {
	Name string `json:"name"`
}

func FromJSONCreatePermissionRequest(body io.Reader) (*CreatePermissionRequest, error) {
	createPermissionRequest := CreatePermissionRequest{}
	if err := json.NewDecoder(body).Decode(&createPermissionRequest); err != nil {
		return nil, err
	}

	if createPermissionRequest.Name == "" {
		return nil, errors.New("'name' is required")
	}

	return &createPermissionRequest, nil
}
