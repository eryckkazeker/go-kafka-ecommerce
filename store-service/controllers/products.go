package controllers

import (
	"net/http"

	fakedatabase "com.ekazeker/store-service/fake-database"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, fakedatabase.GetAllProducts())

}

func GetProductById(c *gin.Context) {

	product, err := fakedatabase.GetProductById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}

	c.JSON(http.StatusOK, product)
}
