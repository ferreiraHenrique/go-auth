package permissionrouter

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/di"
	"github.com/gorilla/mux"
)

func Router(conn sqlite.PoolInterface, r *mux.Router) {
	permissionService := di.ConfigPermissionDI(conn)

	r.Handle("", http.HandlerFunc(permissionService.Create)).Methods("POST")
	r.Handle("", http.HandlerFunc(permissionService.ListByManager)).Methods("GET")

	userService := di.ConfigUserDI(conn)
	r.Use(userService.IsManager)
}
