package model

import (
	"coconut/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string     `json:"title"`
	BodyHTML    *string    `json:"body_html"`
	PublishedAt *time.Time `json:"published_at"`
	Vendor      *string    `json:"vendor"`
	Keywords    *string    `json:"keywords"`
	Price       float32    `json:"price" sql:"default:0.0"`
	Slug        *string    `json:"slug"`
	StockQty    int        `json:"stock_qty" sql:"default:0"`
	Status      int        `json:"status" sql:"default:0"`
	HotSale     bool       `json:"hot_sale" sql:"default:false"`
	NewArrival  bool       `json:"new_arrival"  sql:"default:true"` // 不需要 default:true 否则会有 bug
	CategoryID  int        `json:"category_id"`

	Cover *string
	// Category    Category
}

func GetProducts() []Product {
	var topics []Product
	db.PG.Find(&topics)
	return topics
}

func GetProductById(id string) (Product, error) {
	var tp Product
	if err := db.PG.Where("id = ?", id).First(&tp).Error; err != nil {
		return tp, err
	} else {
		return tp, nil
	}
}

// PRODUCT VALIDATOR

type ProductValidator struct {
	ProductTmp struct {
		Title       string     `form:"title" json:"title" binding:"required"`
		BodyHTML    *string    `form:"body_html" json:"body_html"`
		PublishedAt *time.Time `form:"published_at" json:"published_at"`
		Vendor      *string    `form:"vendor" json:"vendor"`
		Keywords    *string    `form:"keywords" json:"keywords"`
		Price       float32    `form:"price" json:"price"`
		Slug        *string    `form:"slug" json:"slug" binding:"required"`
		StockQty    int        `form:"stock_qty" json:"stock_qty"`
		Status      int        `form:"status" json:"status"`
		HotSale     bool       `form:"hot_sale" json:"hot_sale"`
		NewArrival  bool       `form:"new_arrival" json:"new_arrival"`
		CategoryID  int        `form:"category_id" json:"category_id"`
	} `json:"product"`
	ProductModel Product `json:"-"`
}

func (self *ProductValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(self, b)

	self.ProductModel.Title = self.ProductTmp.Title
	self.ProductModel.BodyHTML = self.ProductTmp.BodyHTML
	self.ProductModel.PublishedAt = self.ProductTmp.PublishedAt
	self.ProductModel.Vendor = self.ProductTmp.Vendor
	self.ProductModel.Keywords = self.ProductTmp.Keywords
	self.ProductModel.Price = self.ProductTmp.Price
	self.ProductModel.Slug = self.ProductTmp.Slug
	self.ProductModel.StockQty = self.ProductTmp.StockQty
	self.ProductModel.Status = self.ProductTmp.Status
	self.ProductModel.HotSale = self.ProductTmp.HotSale
	self.ProductModel.NewArrival = self.ProductTmp.NewArrival
	self.ProductModel.CategoryID = self.ProductTmp.CategoryID

	return err
}
