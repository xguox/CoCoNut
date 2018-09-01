package controller

import (
	"coconut/db"
	"coconut/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	if err := model.SaveData(&category); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   category.ID,
		"name": category.Name,
	})
}

func FetchCategories(c *gin.Context) {
	var categories []model.Category
	db.PG.Find(&categories)

	if len(categories) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no categories found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
}

func UpdateCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "category not found"})
	} else {
		err := c.BindJSON(&category)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
			return
		}

		if err := model.SaveData(&category); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   category.ID,
			"name": category.Name,
		})
	}
}

func DestroyCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "category not found"})
	} else {
		db.PG.Model(&category).Update("DeletedAt", time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}
