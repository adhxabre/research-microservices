package routes

import (
	"users/handlers"

	"github.com/gorilla/mux"
)

func HelloWorldRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HelloWorld).Methods("GET")
}
