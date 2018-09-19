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
	Slug        string     `json:"slug"`
	StockQty    int        `json:"stock_qty" sql:"default:0"`
	Status      int        `json:"status" sql:"default:0"`
	HotSale     bool       `json:"hot_sale" sql:"default:false"`
	NewArrival  bool       `json:"new_arrival"` // 不需要 default:true 否则会有 bug
	CategoryID  int        `json:"category_id"`
	Tags        []Tag      `gorm:"many2many:taggings;"`
	Taggings    []Tagging
	Category    Category
	Cover       *string
	Variants    []Variant
}

func GetProducts() []Product {
	var categories []Product
	db.GetDB().Find(&categories)
	return categories
}

func GetProductByID(id string) (Product, error) {
	var p Product
	tran := db.GetDB().Begin()
	tran.Where("id = ?", id).First(&p)
	tran.Model(&p).Related(&p.Category, "Category")
	tran.Model(p).Related(&p.Tags, "Tags")
	err := tran.Commit().Error
	return p, err
}

func (p *Product) GetCategory() error {
	tran := db.GetDB().Begin()
	tran.Model(p).Related(&p.Category, "Category")
	err := tran.Commit().Error
	return err
}

func (p *Product) GetTags() error {
	tran := db.GetDB().Begin()
	tran.Model(p).Related(&p.Tags, "Tags")
	err := tran.Commit().Error
	return err
}

func (p *Product) SetTag(tagName string) error {
	db := db.GetDB()
	var _t Tag
	var tagging Tagging
	db.FirstOrCreate(&_t, Tag{Name: tagName})

	// pq: null value in column "created_at" violates not-null constraint
	// db.Model(&s).Association("Tag").Append(_t)
	db.Where(Tagging{TagID: _t.ID, ProductID: p.ID}).Attrs(Tagging{CreatedAt: time.Now(), UpdatedAt: time.Now()}).FirstOrCreate(&tagging)
	return nil
}

func (p *Product) RemoveTag(tagName string) {
	db := db.GetDB()
	var _t Tag
	db.Where("name = ?", tagName).First(&_t)
	db.Model(&p).Association("Tags").Delete(_t)
}

// AfterCreate 默认创建一个没有任何 option 的 Variant
func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	var variant Variant
	variant.Price = p.Price
	variant.Stock = 1
	variant.IsDefault = true
	tx.Model(p).Association("Variants").Append(variant)
	return
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
		Slug        string     `form:"slug" json:"slug" binding:"required"`
		StockQty    int        `form:"stock_qty" json:"stock_qty"`
		Status      int        `form:"status" json:"status"`
		HotSale     bool       `form:"hot_sale" json:"hot_sale"`
		NewArrival  bool       `form:"new_arrival" json:"new_arrival"`
		CategoryID  int        `form:"category_id" json:"category_id"`
	} `json:"product"`
	ProductModel Product `json:"-"`
}

func (pv *ProductValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(pv, b)
	if err != nil {
		return err
	}
	pv.ProductModel.Title = pv.ProductTmp.Title
	pv.ProductModel.BodyHTML = pv.ProductTmp.BodyHTML
	pv.ProductModel.PublishedAt = pv.ProductTmp.PublishedAt
	pv.ProductModel.Vendor = pv.ProductTmp.Vendor
	pv.ProductModel.Keywords = pv.ProductTmp.Keywords
	pv.ProductModel.Price = pv.ProductTmp.Price
	pv.ProductModel.Slug = pv.ProductTmp.Slug
	pv.ProductModel.StockQty = pv.ProductTmp.StockQty
	pv.ProductModel.Status = pv.ProductTmp.Status
	pv.ProductModel.HotSale = pv.ProductTmp.HotSale
	pv.ProductModel.NewArrival = pv.ProductTmp.NewArrival
	pv.ProductModel.CategoryID = pv.ProductTmp.CategoryID
	return nil
}
