package model

import (
	"coconut/db"
	"coconut/util"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

type Category struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required,category_nameuniq"`
	Slug string `form:"slug" json:"slug" binding:"required"`
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
	CategoryModel Category `json:"category"`
}

func (self *CategoryValidator) Bind(c *gin.Context) error {
	err := util.CommonBind(c, self)
	// err := c.ShouldBindWith(self.Category, binding.Query)
	if err != nil {
		return err
	}

	return nil
}

// CategoryNameUniq 一大堆参数都抽象不起来啊 = 。 =
func CategoryNameUniq(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var c Category
	category := currentStructOrField.Interface().(Category)
	db.PG.Where("name = ?", category.Name).First(&c)
	return c.Name == ""
}
