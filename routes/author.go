package routes

import (
	"go-api-native/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()
	
	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("/create", authorcontroller.Create).Methods("POST")
}