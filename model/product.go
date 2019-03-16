package model

import (
	"github.com/xguox/coconut/db"
	"errors"
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
	Options     []Option
}

func GetProducts() []Product {
	var products []Product
	db.GetDB().Preload("Category").Preload("Tags").Find(&products)
	return products
}

func GetProductByID(id string) (Product, error) {
	var p Product
	tran := db.GetDB().Begin()
	if err := tran.Where("id = ?", id).First(&p).Error; err != nil {
		tran.Rollback()
		return p, err
	}
	tran.Model(&p).Related(&p.Category, "Category")
	tran.Model(p).Related(&p.Variants, "Variants")
	tran.Model(p).Related(&p.Tags, "Tags")
	tran.Model(p).Related(&p.Options, "Options")
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

// SoftDeleteVaiants 把关联的 Variants 也 delete
func (p *Product) SoftDeleteVaiants() {
	db.GetDB().Table("variants").Where("product_id = ?", p.ID).Delete(&Variant{})
}

func (p *Product) GetDefaultVariant() *Variant {
	var v Variant
	db.GetDB().Find(&v, "product_id = ? AND is_default = ? AND deleted_at IS NULL", p.ID, true)
	return &v
}

// Options & Variants START
//
//

// OptionExists 是否存在 Name: name 的 Option
func (p *Product) OptionExists(name string) bool {
	notFound := db.GetDB().Where("name = ? AND product_id = ?", name, p.ID).First(&Option{}).RecordNotFound()
	return !notFound
}

// AddOption 添加一个新的 Option(Vals只有一个值), 现有的非 default Variants 的对应值改变即可
func (p *Product) AddOption(option Option) error {
	var column string
	if option.Position == 0 || option.Position == 1 {
		// 添加的是第一个 Option, 此前该 Product 的 Options 为空
		err := p.AddOptions([]Option{option})
		return err
	} else if option.Position == 2 {
		column = "option2"
	} else {
		column = "option3"
	}
	tran := db.GetDB().Begin()
	tran.Where("product_id = ? AND is_default = ?", p.ID, true).Delete(&Variant{})
	tran.Model(&p).Association("Options").Append(option)
	tran.Model(&Variant{}).Where("product_id = ? AND is_default = ?", p.ID, false).Update(column, option.ValuesArr()[0])
	err := tran.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteOption 删除一个 Option
func (p *Product) DeleteOption(option *Option) error {
	if len(option.ValuesArr()) > 1 {
		return errors.New("ValuesArr 长度大于1不能删除")
	}

	tran := db.GetDB().Begin()
	tran.Delete(&option)
	tran.Where("product_id = ?", p.ID).Delete(&Variant{})
	err := tran.Commit().Error
	if err != nil {
		return err
	}
	if len(p.Options) == 1 {
		db.GetDB().Unscoped().Model(&Variant{}).Where("product_id = ? AND is_default = ?", p.ID, true).Update("deleted_at", gorm.Expr("NULL"))
	} else {
		p.RebuildVariants()
	}

	return nil
}

// AddOptions 初始化 Options, delete 所有 variants!
func (p *Product) AddOptions(options []Option) error {
	if len(options) < 1 {
		return nil
	}

	tran := db.GetDB().Begin()
	tran.Where("product_id = ?", p.ID).Delete(&Variant{})
	tran.Where("product_id = ?", p.ID).Delete(&Option{})
	tran.Model(&p).Association("Options").Replace(options)

	err := tran.Commit().Error
	if err != nil {
		return err
	}
	p.RebuildVariants()
	return nil
}

// RebuildVariants 根据 Options 生成相应的 Variants, 原有 Variants 删除???
func (p *Product) RebuildVariants() {
	db := db.GetDB()
	var options []Option
	db.Model(&p).Select([]string{"vals"}).Association("Options").Find(&options)

	variants := VariantsBuilding(options)
	//fmt.Println(variants)
	p.Variants = variants
	db.Save(&p)
	//db.Model(&p).Association("Variants").Append(&variants)
}

func (p *Product) FindOptionByID(id string) (*Option, error) {
	var option Option

	if err := db.GetDB().Where("product_id = ? AND id = ?", p.ID, id).First(&option).Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (p *Product) FindVariantByID(id string) (*Variant, error) {
	var variant Variant
	if err := db.GetDB().Where("product_id = ? AND id = ?", p.ID, id).First(&variant).Error; err != nil {
		return nil, err
	}
	return &variant, nil
}

// Options & Variants END

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
