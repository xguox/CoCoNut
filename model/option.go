package model

import (
	"github.com/xguox/coconut/db"
	"github.com/xguox/coconut/util"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Option struct {
	ID        uint `gorm:"primary_key" json:"id" structs:"id"`
	Name      string
	Position  int
	ProductID uint
	Product   Product
	Vals      []byte  // Can't be values, coz MySQL keyword
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Option) ValuesArr() (data []string) {
	json.Unmarshal(o.Vals, &data)
	return
}

func (o *Option) SetValues(data []string) {
	o.Vals, _ = json.Marshal(data)
	return
}

func (o *Option) AddValue(newVal string) error {
	currentValues := o.ValuesArr()

	if util.SliceContains(currentValues, newVal) {
		return nil
	}
	var options []Option
	db := db.GetDB()
	db.Model(&o).Related(&o.Product)
	db.Where("product_id = ?", o.ProductID).Find(&options)
	for i, option := range options {
		if option.Name == o.Name {
			option.SetValues([]string{newVal})
			options[i] = option
			break
		}
	}
	variants := VariantsBuilding(options)
	tran := db.Begin()
	tran.Model(&o).Update("vals", append(currentValues, newVal))
	tran.Model(&o.Product).Association("Variants").Append(variants)
	err := tran.Commit().Error

	return err
}

func (o *Option) RemoveValue(rmVal string) error {
	currentValues := o.ValuesArr()
	var found = false
	for i, val := range currentValues {
		if val == rmVal {
			currentValues = append(currentValues[:i], currentValues[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	var column string
	position := o.Position

	if position == 0 || position == 1 {
		column = "option1"
	} else if position == 2 {
		column = "option2"
	} else {
		column = "option3"
	}
	tran := db.GetDB().Begin()
	tran.Where("product_id = ?", o.ProductID).Where(column+" = ?", rmVal).Delete(&Variant{})
	if len(currentValues) == 0 {
		tran.Delete(&o)
	} else {
		tran.Model(&o).Update("vals", currentValues)
	}
	err := tran.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

func VariantsBuilding(options []Option) (variants []Variant) {
	optionsCount := len(options)
	if optionsCount < 1 {
		return
	}
	// TODO: empty string
	for _, option1 := range util.UniqSlice(options[0].ValuesArr()) {
		if optionsCount > 1 {
			for _, option2 := range util.UniqSlice(options[1].ValuesArr()) {
				if optionsCount > 2 {
					for _, option3 := range util.UniqSlice(options[2].ValuesArr()) {
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
	return
}

// Options Validator
type OptionsValidator struct {
	OptionsTmp []OptionParams `json:"options"`
	Options    []Option       `json:"-"`
}

type OptionParams struct {
	Name   string   `json:"name"`
	Vals []string `json:"vals"`
}

func (o *OptionsValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(o, b)

	if err != nil {
		return err
	}

	var optionsMap = make(map[string][]string)

	for index, option := range o.OptionsTmp {
		optionValues := option.Vals

		if index > 2 {
			break
		}
		if option.Name == "" || len(optionValues) < 1 {
			continue
		}
		optionsMap[option.Name] = append(optionsMap[option.Name], optionValues...)
	}
	initPosition := 0
	for name, values := range optionsMap {
		initPosition++
		_option := Option{Name: name, Position: initPosition}
		_option.SetValues(values)
		o.Options = append(o.Options, _option)
	}
	return nil
}
