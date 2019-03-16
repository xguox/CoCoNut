package controller

import (
	"github.com/xguox/coconut/db"
	"github.com/xguox/coconut/model"
	. "github.com/xguox/coconut/serializer"
	"github.com/xguox/coconut/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductVariants(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	var variants = _product.Variants
	if len(_product.Variants) == 0 {
		variants = append(variants, *_product.GetDefaultVariant())
	}
	s := VariantsSerializer{variants}
	c.JSON(http.StatusOK, gin.H{"data": s.Response()})
}

func ProductVariant(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	_variant, err := _product.FindVariantByID(c.Params.ByName("variant_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no variant found"})
	} else {
		s := VariantSerializer{*_variant}
		c.JSON(http.StatusOK, s.Response())
	}
}

func UpdateVariant(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	_variant, err := _product.FindVariantByID(c.Params.ByName("variant_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no variant found"})
	} else {

		v := model.VariantValidator{VariantModel: *_variant}
		if err := v.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
			return
		}

		db.GetDB().Save(&v.VariantModel)

		s := VariantSerializer{v.VariantModel}
		c.JSON(http.StatusOK, s.Response())
	}
}
