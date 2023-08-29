package permissionservice

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

	permissionRequest, err := dto.FromJSONCreatePermissionRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	permission, err := service.usecase.Create(permissionRequest, manager.ID)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"name": permission.Name,
	})
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(201)
	response.Write(jsonResponse)
}
