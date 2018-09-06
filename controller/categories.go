package controller

import (
	"coconut/model"

	"coconut/util"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	v := model.CategoryValidator{}
	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}

	if err := model.SaveData(&v.CategoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   v.CategoryModel.ID,
		"name": v.CategoryModel.Name,
	})
}

func FetchCategories(c *gin.Context) {
	categories := model.GetCategories()
	if len(categories) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no categories found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
}

func UpdateCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Category not found"})
		return
	}

	v := model.CategoryValidator{CategoryModel: category}

	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}

	if err := model.SaveData(&v.CategoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   v.CategoryModel.ID,
		"name": v.CategoryModel.Name,
	})
}

func DestroyCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "category not found"})
	} else {
		category.SetDeletedAt(time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}
