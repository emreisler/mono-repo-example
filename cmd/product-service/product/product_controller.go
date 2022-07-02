package product

import (
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := getProducts()
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		return err
	}
	return nil
}
