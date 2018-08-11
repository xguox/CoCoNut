package model

import (
	"coconut/db"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	//struct字段之后的tag 因为输出字段的名称默认都是大写的，能够被赋值的字段必须是可导出字段(即首字母大写），同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，要是想通过小写的方式输出 就需要采用json tag的形式
	Name string `json:name`
	Sku  string `json:sku`
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
