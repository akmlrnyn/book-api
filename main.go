package main

import (
	"fmt"
	"go-api-native/config"
	"go-api-native/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	r := mux.NewRouter()


	routes.RouteIndex(r)

	log.Println("Server is now running", config.ENV.DB_PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}