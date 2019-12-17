package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Product .
type Product struct {
	DB *gorm.DB
}

func (db *Product) GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "",
	})
}
