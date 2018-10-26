package model

import (
	"github.com/jinzhu/gorm"
)

type Variant struct {
	gorm.Model
	Price     float32
	Sku       string
	Stock     int
	Position  int
	ProductID uint
	Product   Product
	IsDefault bool
	Option1   string
	Option2   string
	Option3   string
}
