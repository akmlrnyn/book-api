package routes

import "github.com/gorilla/mux"

func RouteIndex(r *mux.Route) {
	api := r.PathPrefix("/api").Subrouter()
	AuthorRoutes(api)
}