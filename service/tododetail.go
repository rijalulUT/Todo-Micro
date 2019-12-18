package service

import (
	"fmt"
	"todomicro/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ToDoDetail struct {
	DB *gorm.DB
}

func (p *ToDoDetail) GetToDoDetail(c *gin.Context) {
	db := p.DB
	var detail model.TodoDetail
	id := c.Param("id")
	db.First(&detail, id)

	c.JSON(200, gin.H{
		"success": true,
		"data":    detail,
	})
}

func (p *ToDoDetail) CreateToDoDetail(c *gin.Context) {
	request := struct {
		Data struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
	}{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(&request)
	var category model.TodoDetail
	category.Title = request.Data.Title
	category.Content = request.Data.Content
	category.CategoryID = 1

	p.DB.Create(&category)
	c.JSON(200, gin.H{
		"success": "true",
		"data":    &category,
	})
}
