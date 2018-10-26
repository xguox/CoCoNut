package model

import (
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
		optionsMap[option.Name] = option.Values
	}
	for name, values := range optionsMap {
		o.Options = append(o.Options, Option{Name: name, Values: values})
	}
	return nil
}

// Option Validator

type OptionValidator struct {
	OptionTmp struct {
		Name   string         `form:"name" json:"name" binding:"required"`
		Values pq.StringArray `form:"values" json:"values"`
	} `json:"option"`
	OptionModel Option `json:"-"`
}

func (ov *OptionValidator) Bind(c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(ov, b)
	if err != nil {
		return err
	}
	ov.OptionModel.Name = ov.OptionTmp.Name
	ov.OptionModel.Values = ov.OptionTmp.Values
	return nil
}
