package clientrouter

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/di"
	"github.com/gorilla/mux"
)

func Router(conn sqlite.PoolInterface, r *mux.Router) {
	clientService := di.ConfigClientDI(conn)

	r.Handle("", http.HandlerFunc(clientService.Create)).Methods("POST")
}
