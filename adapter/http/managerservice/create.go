package managerservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	managerRequest, err := dto.FromJSONCreateManagerRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	manager, err := service.usecase.Create(managerRequest)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"username": manager.User.Username,
	})
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
