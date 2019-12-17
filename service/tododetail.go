package service

import (
	"todomicro/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CategoryDetail struct {
	DB *gorm.DB
}

type DetailCategoryList struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId int32  `json:"category_id"`
}

func (p *CategoryDetail) GetDetailCategory(c *gin.Context) {
	db := p.DB
	var detail model.TodoDetail
	id := c.Param("id")
	db.First(&detail, id)

	c.JSON(200, gin.H{
		"success": true,
		"data":    detail,
	})
}
