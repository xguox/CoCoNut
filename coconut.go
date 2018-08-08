package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	//	open a db connection
	var err error
	db, err = gorm.Open("postgres", "user=postgres dbname=coconut_development sslmode=disable")
	if err != nil {
		panic("failed to connect db")
	}
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/products")
	{
		v1.POST("/", createProduct)
		v1.GET("/", fetchAllProducts)
		v1.GET("/:id", fetchProduct)
		v1.PUT("/:id", updateProduct)
		v1.DELETE("/:id", destroyProduct)
	}
	router.Run(":9876")
}

type (
	product struct {
		gorm.Model
		//struct字段之后的tag 因为输出字段的名称默认都是大写的，能够被赋值的字段必须是可导出字段(即首字母大写），同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，要是想通过小写的方式输出 就需要采用json tag的形式
		Name string `json:title`
		Sku  string `json:completed`
	}
)

func createProduct(c *gin.Context) {
	_product := product{Name: c.PostForm("name"), Sku: c.PostForm("sku")}
	db.Save(&_product)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Product created successfully!", "resourceId": _product.ID})
}

func fetchAllProducts(c *gin.Context) {
	var products []product

	db.Find(&products)

	if len(products) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": products})
}

func fetchProduct(c *gin.Context) {
	var _product product
	id := c.Params.ByName("id")
	if db.First(&_product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		c.JSON(http.StatusOK, _product)
	}
}

func updateProduct(c *gin.Context) {
	var _product product
	id := c.Params.ByName("id")
	if db.First(&_product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		_product.Name = c.PostForm("name")
		_product.Sku = c.PostForm("sku")
		db.Save(&_product)
		c.JSON(http.StatusOK, _product)
	}
}

func destroyProduct(c *gin.Context) {
	var _product product
	id := c.Params.ByName("id")

	if db.First(&_product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
	} else {
		db.Model(&_product).Update("DeletedAt", time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}
