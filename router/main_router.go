package router

import (
	"fmt"
	"net/http"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func (rh RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := rh(w, r) // Call handler function
	if err != nil {
		fmt.Println(err)
	}
}
