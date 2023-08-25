package userrouter

import (
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/di"
	"github.com/gorilla/mux"
)

func Router(conn sqlite.PoolInterface, r *mux.Router) {
	userService := di.ConfigUserDI(conn)

	r.Handle("/signin", http.HandlerFunc(userService.Signin)).Methods("POST")
}
