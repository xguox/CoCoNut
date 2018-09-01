package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug`
}
