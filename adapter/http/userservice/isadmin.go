package userservice

import (
	"net/http"
)

func (service service) IsAdmin(h http.Handler) http.Handler {
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

		if !service.usecase.IsAdmin(auth[0]) {
			response.WriteHeader(403)
			response.Write([]byte("missing admin authorization"))
			return
		}

		h.ServeHTTP(response, request)
	})
}
