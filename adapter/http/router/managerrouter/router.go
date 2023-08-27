package managerrouter

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/di"
	"github.com/gorilla/mux"
)

func Router(conn sqlite.PoolInterface, r *mux.Router) {
	managerService := di.ConfigManagerDI(conn)

	r.Handle("", http.HandlerFunc(managerService.Create)).Methods("POST")

	userService := di.ConfigUserDI(conn)
	r.Use(userService.IsAdmin)
}
