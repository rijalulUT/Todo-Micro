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
	detail := service.CategoryDetail{DB: db}

	r.GET("/category", category.GetCategory)
	r.GET("/detail/:id", detail.GetDetailCategory)
	r.POST("/category", category.CreateCategory)
	r.DELETE("/category/:id", category.DeleteCategory)
	r.PUT("/category/:id", category.UpdateCategory)

	r.Run()
}
