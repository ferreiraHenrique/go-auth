package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (service service) Signin(response http.ResponseWriter, request *http.Request) {
	userRequest, err := dto.FromJSONSigninUserRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usercase.Signin(userRequest)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"username": user.Username,
	})
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
