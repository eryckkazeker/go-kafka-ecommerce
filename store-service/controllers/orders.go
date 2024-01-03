package controllers

import (
	"fmt"
	"net/http"

	fakedatabase "com.ekazeker/store-service/fake-database"
	"com.ekazeker/store-service/models"
	"com.ekazeker/store-service/services"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(ctx *gin.Context) {

	var stock []models.ProductStock

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order := models.NewOrder()

	for _, prod := range stock {

		if qtd := fakedatabase.GetProductStockById(prod.Id); qtd < prod.Quantity {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%s not in stock", prod.Id),
			})

			return
		}

		order.Products = append(order.Products, prod)

	}

	services.CalculatePrice(order)

	services.SendOrder(*order)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order Placed. Id: %s", order.Id),
	})

}
