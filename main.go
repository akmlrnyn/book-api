package main

import (
	"fmt"
	"go-api-native/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	r := mux.NewRouter()

	log.Println("Server is now running ")
	http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r)
}