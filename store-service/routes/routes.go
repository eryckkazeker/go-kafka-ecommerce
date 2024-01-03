package routes

import (
	"com.ekazeker/store-service/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.POST("/orders", controllers.PlaceOrder)
	r.Run()
}
