package clientservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	clientRequest, err := dto.FromJSONCreateClientRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	client, err := service.usecaase.Create(clientRequest)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"username": client.User.Username,
	})
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
