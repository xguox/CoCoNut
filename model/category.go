package model

import (
	"coconut/db"
	"fmt"
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required,is-uniq"`
	Slug string `form:"slug" json:"slug" binding:"required"`
}

func GetCategoryByID(id string) (Category, error) {
	var category Category
	if err := db.GetDB().Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func GetCategories() []Category {
	var categories []Category
	db.GetDB().Find(&categories)
	return categories
}

func (c *Category) SetDeletedAt(t time.Time) {
	db.GetDB().Model(c).Update("DeletedAt", time.Now())
}

// CATEGORY VALIDATOR

type CategoryValidator struct {
	CategoryModel Category `json:"category"`
}

func (s *CategoryValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}
	return nil
}

func ValidateUniq(fl validator.FieldLevel) bool {
	var result struct{ Count int }
	currentField, _, _ := fl.GetStructFieldOK()
	table := modelTableNameMap[currentField.Type().Name()] // table name
	value := fl.Field().String()                           // value
	column := fl.FieldName()                               // column name
	sql := fmt.Sprintf("select count(*) from %s where %s='%s'", table, column, value)
	db.GetDB().Raw(sql).Scan(&result)
	dup := result.Count > 0
	return !dup
}
