package clientservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
	"github.com/gorilla/context"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	manager := context.Get(request, "manager").(*domain.Manager)
	if manager == nil {
		response.WriteHeader(400)
		response.Write([]byte("manager not found"))
		return
	}

	clientRequest, err := dto.FromJSONCreateClientRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	client, err := service.usecaase.Create(clientRequest, manager.ID)
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
