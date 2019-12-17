package main

import (
	"todomicro/config"
	"todomicro/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	product := service.Product{DB: db}

	r.GET("/products", product.GetProducts)

	r.Run()
}
