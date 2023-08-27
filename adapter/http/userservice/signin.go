package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (service service) Signin(response http.ResponseWriter, request *http.Request) {
	userRequest, err := dto.FromJSONSigninUserRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.Signin(userRequest)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	token := domain.NewToken()
	token.SetClaim("username", user.Username)
	tokenSigned, _ := token.SignString()

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
		"token":    tokenSigned,
	})
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
