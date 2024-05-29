package routes

import (
	"go-api-native/controllers/bookcontroller"
	"go-api-native/middleware"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.Use(middleware.Auth)

	router.HandleFunc("", bookcontroller.Index)
	router.HandleFunc("/create", bookcontroller.Create)
	router.HandleFunc("/{id}/detail", bookcontroller.Detail)
	router.HandleFunc("/{id}/update", bookcontroller.Update)
	router.HandleFunc("/{id}/delete", bookcontroller.Delete)
}