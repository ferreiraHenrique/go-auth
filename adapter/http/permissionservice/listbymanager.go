package permissionservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/gorilla/context"
)

type PermissionResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type JSONResponse struct {
	Permissions []PermissionResponse `json:"permissions"`
}

func (service service) ListByManager(response http.ResponseWriter, request *http.Request) {
	manager := context.Get(request, "manager").(*domain.Manager)
	if manager == nil {
		response.WriteHeader(400)
		response.Write([]byte("manager not found"))
		return
	}

	permissions, err := service.usecase.ListByManager(manager.ID)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	if len(*permissions) == 0 {
		response.WriteHeader(404)
		response.Write([]byte("permissions not found"))
		return
	}

	resp := JSONResponse{}
	for _, p := range *permissions {
		permissionResponse := PermissionResponse{
			UUID: p.UUID,
			Name: p.Name,
		}
		resp.Permissions = append(resp.Permissions, permissionResponse)
	}

	jsonResponse, _ := json.Marshal(resp)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
