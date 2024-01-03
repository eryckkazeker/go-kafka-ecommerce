package fakedatabase

import (
	"errors"

	"com.ekazeker/store-service/models"
)

var productMap = map[string]models.Product{
	"p1": {Id: "p1", Description: "First Product", Price: 12350},
	"p2": {Id: "p2", Description: "Second Product", Price: 2290},
}

func GetAllProducts() []models.Product {

	prodArray := make([]models.Product, 0)

	for _, v := range productMap {
		prodArray = append(prodArray, v)
	}

	return prodArray
}

func GetProductById(id string) (models.Product, error) {
	product, ok := productMap[id]

	if !ok {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}
