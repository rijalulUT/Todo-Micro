package service

import (
	"fmt"
	"todomicro/model"

	"strconv"

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
	if len(category) == 0 {
		c.JSON(200, gin.H{
			"success": false,
			"data":    category,
		})
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    category,
	})
}

func (p *Category) CreateCategory(c *gin.Context) {
	request := struct {
		Data struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(&request)
	var category model.TodoCategory
	category.Title = request.Data.Title
	category.Description = request.Data.Description
	category.UserID = 1

	p.DB.Create(&category)
	c.JSON(200, gin.H{
		"success": "true",
		"data":    &category,
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

// Delete Dzakiy
func (p *Category) DeleteCategory(c *gin.Context) {
	db := p.DB
	var category model.TodoCategory
	var detail model.TodoDetail

	id := c.Param("id")
	uintid, _ := strconv.ParseUint(id, 10, 64)
	db.Delete(&category, id)
	db.Delete(&detail, model.TodoDetail{
		CategoryID: uintid,
	})

	c.JSON(200, gin.H{
		"success": true,
		"data":    category,
	})
}

func (p *Category) UpdateCategory(c *gin.Context) {
	db := p.DB
	var category model.TodoCategory

	request := struct {
		Data struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")
	uintid, _ := strconv.ParseUint(id, 10, 64)
	category.ID = uintid

	db.First(&category)
	category.Title = request.Data.Title
	category.Description = request.Data.Description

	db.Save(&category)

	c.JSON(500, gin.H{
		"success": "true",
		"data":    &category,
	})

}

func (p *Category) UpdateDetail(c *gin.Context) {
	db := p.DB
	var detail model.TodoDetail

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

	id := c.Param("id")
	uintid, _ := strconv.ParseUint(id, 10, 64)
	detail.ID = uintid

	db.First(&detail)
	detail.Title = request.Data.Title
	detail.Content = request.Data.Content

	db.Save(&detail)

	c.JSON(500, gin.H{
		"success": "true",
		"data":    &detail,
	})

}
