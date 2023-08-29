package clientservice

import (
	"encoding/json"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/gorilla/context"
)

type ClientResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type JSONResponse struct {
	Clients []ClientResponse `json:"clients"`
}

func (service service) ListByManager(response http.ResponseWriter, request *http.Request) {
	manager := context.Get(request, "manager").(*domain.Manager)
	if manager == nil {
		response.WriteHeader(400)
		response.Write([]byte("manager not found"))
		return
	}

	clients, err := service.usecase.ListByManager(manager.ID)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	if len(*clients) == 0 {
		response.WriteHeader(404)
		response.Write([]byte("clients not found"))
		return
	}

	resp := JSONResponse{}
	for _, c := range *clients {
		clientRespose := ClientResponse{
			UUID: c.UUID,
			Name: c.Name,
		}
		resp.Clients = append(resp.Clients, clientRespose)
	}

	jsonResponse, _ := json.Marshal(resp)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonResponse)
}
