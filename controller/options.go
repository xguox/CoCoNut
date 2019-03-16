package controller

import (
	"github.com/xguox/coconut/model"
	. "github.com/xguox/coconut/serializer"
	"github.com/xguox/coconut/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchOptions 列出某个 Product 的 Options
func FetchOptions(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}

	s := OptionsSerializer{_product.Options}
	c.JSON(http.StatusOK, gin.H{"data": s.Response()})

}

// InitBuildOptions 仅当 IsDefault 为 true 的 Variant 的 deleted_at 不为空的时候初始化 options 组合
func InitBuildOptions(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	defaultVariant := _product.GetDefaultVariant()

	if defaultVariant == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "Oops!"})
		return
	}

	v := model.OptionsValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}
	if err = _product.AddOptions(v.Options); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Options created successfully!"})
}

// CreateOption 已存在一个或多个 Options 组合时候, 添加新的 option, 新添加的 option 只能有一个 value
func CreateOption(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	var reqJSON struct {
		Name  string
		Value string
	}
	c.BindJSON(&reqJSON)
	if reqJSON.Name == "" || reqJSON.Value == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "无效参数"})
		return
	}
	if _product.OptionExists(reqJSON.Name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Option 已存在"})
		return
	}
	position := len(_product.Options) + 1
	if position > 3 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Option 最多3个"})
		return
	}

	_option := model.Option{Name: reqJSON.Name, Position: position}
	_option.SetValues([]string{reqJSON.Value})
	err = _product.AddOption(_option)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Option created successfully!"})
	}
}

// AddSingleValue 单独给一个 option 加一个 value
func AddSingleValue(c *gin.Context) {
	_product, err := model.GetProductByID(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	option, err := _product.FindOptionByID(c.Params.ByName("option_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no option found"})
		return
	}
	var reqJSON struct {
		Value string
	}

	c.BindJSON(&reqJSON)
	if reqJSON.Value == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "无效参数"})
		return
	}
	if err = option.AddValue(reqJSON.Value); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value add successfully!"})
}

// DeleteOption 删除单个 option (仅当 option 只有一个 value 时候可以操作)
func DeleteOption(c *gin.Context) {
	_product, err := model.GetProductByID(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	option, err := _product.FindOptionByID(c.Params.ByName("option_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no option found"})
		return
	}

	if err = _product.DeleteOption(option); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已删除!"})
}

// DeleteSingleValue 删除单个 option 的单个 value
func DeleteSingleValue(c *gin.Context) {
	_product, err := model.GetProductByID(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	option, err := _product.FindOptionByID(c.Params.ByName("option_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no option found"})
		return
	}

	var reqJSON struct {
		Value string
	}

	c.BindJSON(&reqJSON)
	if err = option.RemoveValue(reqJSON.Value); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value remove successfully!"})
}

// ReorderOptions 排列 Options
func ReorderOptions(c *gin.Context) {

}
