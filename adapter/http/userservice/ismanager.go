package userservice

import (
	"net/http"
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

		if !service.usecase.IsManager(auth[0]) {
			response.WriteHeader(403)
			response.Write([]byte("missing manager authorization"))
			return
		}

		h.ServeHTTP(response, request)
	})
}
