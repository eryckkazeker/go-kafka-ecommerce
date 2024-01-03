package services

import (
	fakedatabase "com.ekazeker/store-service/fake-database"
	"com.ekazeker/store-service/models"
)

func CalculatePrice(order *models.Order) {
	totalAmount := 0
	for _, productStock := range order.Products {
		product, _ := fakedatabase.GetProductById(productStock.Id)
		totalAmount += product.Price
	}
	order.TotalAmount = totalAmount
}
