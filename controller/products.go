package controller

import (
	"coconut/db"
	"coconut/model"
	. "coconut/serializer"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	_product := model.CreateProduct(c.PostForm("name"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product created successfully!", "resourceId": _product.ID})
}

func FetchAllProducts(c *gin.Context) {
	products := model.GetProducts()
	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	s := ProductsSerializer{products}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s.Response()})
}

func FetchProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		s := ProductSerializer{_product}
		c.JSON(http.StatusOK, s.Response())
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		_product.Update(c.PostForm("name"))
		s := ProductSerializer{_product}
		c.JSON(http.StatusOK, s.Response())
	}
}

func DestroyProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		db.PG.Model(&_product).Update("DeletedAt", time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}
