package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferreiraHenrique/go-auth/adapter/http/router/managerrouter"
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

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
