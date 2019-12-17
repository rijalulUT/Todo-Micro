package service

import (
	"todomicro/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Category struct {
	DB *gorm.DB
}

type Hasil struct {
	Success  bool       `json:"success"`
	Category []Category `json:"category"`
}
type CategoryList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Userid      int32  `json:"user_id"`
}

func (p *Category) GetCategory(c *gin.Context) {
	db := p.DB
	var category []model.TodoCategory
	// var Categorys []model.Category

	db.Find(&category)

	c.JSON(200, gin.H{
		"success": true,
		"data":    category,
	})
}

// func (p *Product) CreateProduct(c *gin.Context) {
// 	var request product

// 	if err := c.ShouldBind(&request); err != nil {
// 		c.JSON(500, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	p.DB.Create(&request)
// 	c.JSON(200, gin.H{
// 		"message": "success",
// 	})
// }
