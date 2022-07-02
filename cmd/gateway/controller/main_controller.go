package controller

import (
	"encoding/json"
	"example.com/gateway/client"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) error {
	userClient := client.GetUserClient()
	users, err := userClient.GetUsers()
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return err
	}
	return nil
}

func GetProducts(w http.ResponseWriter, r *http.Request) error {
	productClient := client.GetProductClient()
	products, err := productClient.GetProducts()
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		return err
	}
	return nil
}
