package controller

import (
	"coconut/db"
	"coconut/model"
	. "coconut/serializer"
	"coconut/util"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	v := model.ProductValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}
	if checkCategoryNotFound(c, v.ProductTmp.CategoryID) {
		return
	}
	if err := model.SaveData(&v.ProductModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	s := ProductSerializer{v.ProductModel}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully!", "data": s.Response()})
}

func FetchAllProducts(c *gin.Context) {
	products := model.GetProducts()
	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	s := ProductsSerializer{products}

	c.JSON(http.StatusOK, gin.H{"data": s.Response()})
}

func FetchProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		s := ProductSerializer{_product}
		c.JSON(http.StatusOK, s.Response())
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := model.GetProductByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		v := model.ProductValidator{ProductModel: product}
		if err := v.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
			return
		}
		if checkCategoryNotFound(c, v.ProductTmp.CategoryID) {
			return
		}

		db.PG.Save(&v.ProductModel)
		s := ProductSerializer{v.ProductModel}
		c.JSON(http.StatusOK, s.Response())
	}
}

func DestroyProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		db.PG.Model(&_product).Update("DeletedAt", time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}

func checkCategoryNotFound(c *gin.Context, id int) bool {
	_, err := model.GetCategoryByID(fmt.Sprintf("%d", id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "category not found"})
		return true
	}
	return false
}
