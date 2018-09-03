package model

import (
	"coconut/db"
	"reflect"

	"github.com/gin-gonic/gin/binding"

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
	// err := util.CommonBind(c, self)
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(self, b)
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
	db.PG.Where("name = ? AND id != ?", category.Name, category.ID).First(&c)
	return c.Name == ""
}
