package model

import (
	"github.com/xguox/coconut/db"
	"fmt"
	"reflect"
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
	Slug string
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
	CategoryTmp struct {
		Name string `form:"name" json:"name" binding:"required,is-uniq"`
		Slug string `form:"slug" json:"slug" binding:"required"`
	} `json:"category"`
	CategoryModel Category `json:"-"`
}

func (c *CategoryValidator) TableName() string {
	return "categories"
}

func (s *CategoryValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(s, b)
	if err != nil {
		return err
	}

	s.CategoryModel.Name = s.CategoryTmp.Name
	s.CategoryModel.Slug = s.CategoryTmp.Slug
	return nil
}

func ValidateUniq(fl validator.FieldLevel) bool {
	var result struct{ Count int }
	table := fl.Top().MethodByName("TableName").Call([]reflect.Value{})[0]
	value := fl.Field().String() // value
	column := fl.FieldName()     // column name

	sql := fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE %s='%s'", table, column, value)
	db.GetDB().Raw(sql).Scan(&result)
	dup := result.Count > 0
	return !dup
}
