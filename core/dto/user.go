package dto

import (
	"encoding/json"
	"errors"
	"io"
)

type SigninUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromJSONSigninUserRequest(body io.Reader) (*SigninUserRequest, error) {
	signinUserRequest := SigninUserRequest{}
	if err := json.NewDecoder(body).Decode(&signinUserRequest); err != nil {
		return nil, err
	}

	if signinUserRequest.Username == "" {
		return nil, errors.New("'username' is required")
	}

	if signinUserRequest.Password == "" {
		return nil, errors.New("'password' is required")
	}

	return &signinUserRequest, nil
}
