package client

import (
	"encoding/json"
	"example.com/models"
	"fmt"
	"net/http"
)

type ProductClient struct {
	Url string
}

func GetProductClient() *ProductClient {
	return &ProductClient{
		Url: "http://localhost:8084",
	}
}

func (p *ProductClient) GetProducts() ([]models.Product, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", p.Url), nil)
	if err != nil {
		return nil, err
	}
	var products []models.Product
	resp, err := http.DefaultClient.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
