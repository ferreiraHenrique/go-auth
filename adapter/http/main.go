package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/http/router/clientrouter"
	"github.com/ferreiraHenrique/go-auth/adapter/http/router/managerrouter"
	"github.com/ferreiraHenrique/go-auth/adapter/http/router/userrouter"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	conn := sqlite.GetConnection()

	router := mux.NewRouter()

	managerRouter := router.PathPrefix("/manager").Subrouter()
	managerrouter.Router(conn, managerRouter)

	userRouter := router.PathPrefix("/user").Subrouter()
	userrouter.Router(conn, userRouter)

	clientRouter := router.PathPrefix("/client").Subrouter()
	clientrouter.Router(conn, clientRouter)

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
