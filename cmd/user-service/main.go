package main

import (
	"example.com/router"
	"example.com/user-service/user"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/users", router.RootHandler(user.GetUsers))
	http.ListenAndServe(":8083", r)
}
