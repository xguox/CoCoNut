package controller

import (
	"coconut/db"
	"coconut/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	_product := model.CreateProduct(c.PostForm("name"), c.PostForm("sku"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product created successfully!", "resourceId": _product.ID})
}

func FetchAllProducts(c *gin.Context) {
	products := model.GetProducts()
	log.Println("测试 .realize.yaml mark II")
	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": products})
}

func FetchProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		c.JSON(http.StatusOK, _product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		_product.Update(c.PostForm("name"), c.PostForm("sku"))
		c.JSON(http.StatusOK, _product)
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
