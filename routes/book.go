package routes

import (
	"go-api-native/controllers/bookcontroller"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", bookcontroller.Index)
}