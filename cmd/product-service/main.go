package main

import (
	"example.com/product-service/product"
	"example.com/router"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/products", router.RootHandler(product.GetProducts)).Methods("GET")
	http.ListenAndServe(":8084", r)
}
