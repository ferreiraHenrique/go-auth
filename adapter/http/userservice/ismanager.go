package userservice

import (
	"net/http"

	"github.com/gorilla/context"
)

func (service service) IsManager(h http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		auth, ok := request.Header["Authorization"]
		if !ok {
			response.WriteHeader(403)
			response.Write([]byte("missing authorization"))
			return
		}

		if auth[0] == "" {
			response.WriteHeader(403)
			response.Write([]byte("missing authorization"))
			return
		}

		isAdmin, manager := service.usecase.IsManager(auth[0])
		if !isAdmin {
			response.WriteHeader(403)
			response.Write([]byte("missing manager authorization"))
			return
		}

		context.Set(request, "manager", manager)

		h.ServeHTTP(response, request)
	})
}
