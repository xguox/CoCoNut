package model

import (
	"coconut/db"
	"coconut/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetCategoryById(id string) (Category, error) {
	var category Category
	if err := db.PG.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	} else {
		return category, nil
	}
}

// CATEGORY VALIDATOR

type CategoryValidator struct {
	Category struct {
		Name string `form:"name" json:"name" binding:"required"`
		Slug string `form:"slug" json:"slug" binding:"required"`
	} `json:"category"`
	CategoryModel Category `json:"-"`
}

func (self *CategoryValidator) Bind(c *gin.Context) error {
	err := util.CommonBind(c, self)
	// err := c.ShouldBindWith(self.Category, binding.Query)
	if err != nil {
		return err
	}
	self.CategoryModel.Name = self.Category.Name
	self.CategoryModel.Slug = self.Category.Slug
	return nil
}
