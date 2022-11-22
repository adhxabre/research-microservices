package routes

import (
	"users/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", handlers.FindUsers).Methods("GET")
	r.HandleFunc("/user/{index}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{index}", handlers.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{index}", handlers.DeleteUser).Methods("DELETE")
}
