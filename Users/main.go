package main

import (
	"fmt"
	"net/http"
	"users/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Server is Running on localhost:5000")
	http.ListenAndServe(":5000", r)
}
