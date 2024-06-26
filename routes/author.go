package routes

import (
	"go-api-native/controllers/authcontroller"
	"go-api-native/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()
	
	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	// router.HandleFunc("/create", authorcontroller.Create).Methods("POST")
	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/{id}/detail", authorcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", authorcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", authorcontroller.Delete).Methods("DELETE")
}