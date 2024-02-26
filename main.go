package main

import (
	"crud_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/products", getProducts)

	router.GET("/product/:productId", getProduct)

	router.POST("/products", addProduct)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {

	products := models.GetProducts()

	if products == nil || len(products) == 0 {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, products)

	}
}

func getProduct(c *gin.Context) {

	id := c.Param("id")

	//converting string to number
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product := models.GetProduct(productID)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, product)

	}

}

func addProduct(c *gin.Context) {

	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}

}
