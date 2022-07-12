package product

import (
	"example.com/models"
)

//GetProducts Getting products from DB - mock data
func getProducts() ([]models.Product, error) {
	return []models.Product{
		models.Product{1, "Product1", 5, 2, "Apple"},
		models.Product{2, "Product2", 15, 1, "Banana"},
	}, nil
}
