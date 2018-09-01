package model

import (
	"coconut/db"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required"`
	Slug string `json:"slug`
}

func GetCategoryById(id string) (Category, error) {
	var category Category
	if err := db.PG.Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	} else {
		return category, nil
	}
}
