package permissionservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/gorilla/context"
)

type JSONResponse struct {
	Permissions []string `json:"permissions"`
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

	resp := JSONResponse{}
	for _, p := range *permissions {
		resp.Permissions = append(resp.Permissions, p.Name)
	}

	jsonResponse, _ := json.Marshal(resp)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
