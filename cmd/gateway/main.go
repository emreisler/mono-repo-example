package main

import (
	"example.com/gateway/controller"
	"example.com/router"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/users", router.RootHandler(controller.GetUsers)).Methods("GET")
	r.Handle("/products", router.RootHandler(controller.GetProducts)).Methods("GET")
	http.ListenAndServe(":8081", r)
}
