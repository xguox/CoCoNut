package model

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

// VARIANT VALIDATOR

type VariantValidator struct {
	VariantTmp struct {
		Price    float32 `form:"price" json:"price"`
		Stock    int     `form:"stock" json:"stock"`
		Sku      string  `form:"sku" json:"sku" binding:"required"`
		Position int     `form:"position" json:"position"`
	} `json:"variant"`
	VariantModel Variant `json:"-"`
}

func (vv *VariantValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(vv, b)
	if err != nil {
		return err
	}
	vv.VariantModel.Price = vv.VariantTmp.Price
	vv.VariantModel.Stock = vv.VariantTmp.Stock
	vv.VariantModel.Sku = vv.VariantTmp.Sku
	vv.VariantModel.Position = vv.VariantTmp.Position

	return nil
}
