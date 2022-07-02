package client

import (
	"example.com/models"
	"fmt"
	"net/http"
)

type ProductClient struct {
	Url string
}

func GetProductClient() *ProductClient {
	return &ProductClient{
		Url: "localhost:8084",
	}
}

func (p *ProductClient) GetProducts() ([]models.Product, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products", p.Url), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	return nil, nil
}
