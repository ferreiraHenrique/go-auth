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
	r.Handle("", http.HandlerFunc(clientService.ListByManager)).Methods("GET")
	r.Handle("/attach", http.HandlerFunc(clientService.AttachPermission)).Methods("POST")

	userService := di.ConfigUserDI(conn)
	r.Use(userService.IsManager)
}
