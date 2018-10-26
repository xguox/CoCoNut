package controller

import (
	"coconut/model"
	"coconut/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// BuildOptions 已存在一个或多个 Options 组合时候, 添加新的 option, 新添加的 option 只能有一个 value
func BuildOptions(c *gin.Context) {

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
	option.AddValue(reqJSON.Value)
	// TODO: err handle
	c.JSON(http.StatusOK, gin.H{"message": "Value add successfully!"})
}

// DeleteSingleValue 删除单个 option 的单个 value
func DeleteSingleValue(c *gin.Context) {

}

// DeleteOption 删除单个 option (仅当 option 只有一个 value 时候可以操作)
func DeleteOption(c *gin.Context) {

}

// ReorderOptions 排列 Options
func ReorderOptions(c *gin.Context) {

}
