package clientservice

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/dto"
	"github.com/gorilla/context"
)

func (service service) AttachPermission(response http.ResponseWriter, request *http.Request) {
	manager := context.Get(request, "manager").(*domain.Manager)
	if manager == nil {
		response.WriteHeader(400)
		response.Write([]byte("manager not found"))
		return
	}

	clientRequest, err := dto.FromJSONAttachClientPermissionRequest(request.Body)
	if err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	if err := service.usecase.AttachPermission(clientRequest); err != nil {
		response.WriteHeader(400)
		response.Write([]byte(err.Error()))
		return
	}

	response.WriteHeader(204)
}
