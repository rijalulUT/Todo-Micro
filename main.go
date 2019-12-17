package main

import (
	"library/config"
	"library/product"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	product := product.Product{DB: db}

	r.GET("/products", product.GetProducts)
	r.POST("/products", product.CreateProduct)

	r.Run()
}
