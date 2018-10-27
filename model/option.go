package model

import (
	"coconut/db"
	"coconut/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lib/pq"
)

type Option struct {
	ID        uint `gorm:"primary_key" json:"id" structs:"id"`
	Name      string
	Position  int
	ProductID uint
	Product   Product
	Values    pq.StringArray `gorm:"type:varchar(100)[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Option) AddValue(newVal string) error {
	currentValues := o.Values
	for _, val := range currentValues {
		if newVal == val {
			return nil
		}
	}
	var options []Option
	db := db.GetDB()
	db.Model(&o).Related(&o.Product)
	db.Where("product_id = ?", o.ProductID).Find(&options)
	for i, option := range options {
		if option.Name == o.Name {
			option.Values = []string{newVal}
			options[i] = option
			break
		}
	}
	variants := VariantsBuilding(options)
	tran := db.Begin()
	tran.Model(&o).Update("values", append(currentValues, newVal))
	tran.Model(&o.Product).Association("Variants").Append(variants)
	err := tran.Commit().Error

	return err
}

func VariantsBuilding(options []Option) []Variant {
	var variants []Variant

	optionsCount := len(options)

	for _, option1 := range util.UniqSlice(options[0].Values) {
		if optionsCount > 1 {
			for _, option2 := range util.UniqSlice(options[1].Values) {
				if optionsCount > 2 {
					for _, option3 := range util.UniqSlice(options[2].Values) {
						// create variant with 3 options
						variants = append(variants, Variant{Option1: option1, Option2: option2, Option3: option3})
					}
				} else {
					// create variant with 2 options
					variants = append(variants, Variant{Option1: option1, Option2: option2})
				}
			}
		} else {
			// create variant with 1 options
			variants = append(variants, Variant{Option1: option1})
		}
	}
	return variants
}

// Options Validator
type OptionsValidator struct {
	OptionsTmp []OptionParams `json:"options"`
	Options    []Option       `json:"-"`
}

type OptionParams struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

func (o *OptionsValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(o, b)

	if err != nil {
		return err
	}

	var optionsMap = make(map[string][]string)

	for index, option := range o.OptionsTmp {
		if index > 2 {
			break
		}
		if option.Name == "" || len(option.Values) < 1 {
			continue
		}
		optionsMap[option.Name] = append(optionsMap[option.Name], option.Values...)
	}
	initPosition := 0
	for name, values := range optionsMap {
		initPosition++
		o.Options = append(o.Options, Option{Name: name, Values: values, Position: initPosition})
	}
	return nil
}

// Option Validator
//
// type OptionValidator struct {
// 	OptionTmp struct {
// 		Name   string         `form:"name" json:"name" binding:"required"`
// 		Values pq.StringArray `form:"values" json:"values"`
// 	} `json:"option"`
// 	OptionModel Option `json:"-"`
// }

// func (o *OptionValidator) TableName() string {
// 	return "options"
// }

// func (ov *OptionValidator) Bind(c *gin.Context) error {
// 	b := binding.Default(c.Request.Method, c.ContentType())
// 	err := c.ShouldBindWith(ov, b)
// 	if err != nil {
// 		return err
// 	}
// 	ov.OptionModel.Name = ov.OptionTmp.Name
// 	ov.OptionModel.Values = ov.OptionTmp.Values
// 	return nil
// }
