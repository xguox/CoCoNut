package model

import (
	"coconut/db"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	BodyHTML    *string
	PublishedAt *time.Time
	Vendor      *string
	Keywords    *string
	Price       float32
	Slug        string
	StockQty    int
	Status      int
	HotSale     bool
	NewArrival  bool
	Cover       string
	Category    Category
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

func CreateProduct(name, sku string) Product {
	tp := Product{Name: name, Sku: sku}
	db.PG.Create(&tp)
	return tp
}

// TODO: 所有字段都能传参更改, 现在只能改 name, sku,
func (p *Product) Update(name, sku string) {
	p.Name = name
	p.Sku = sku
	db.PG.Save(&p)
}
