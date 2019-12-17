package main

import (
	"todomicro/config"
	"todomicro/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	category := service.Category{DB: db}

	r.GET("/category", category.GetCategory)

	r.Run()
}
