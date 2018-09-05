package main

import (
	. "coconut/controller"

	"github.com/gin-gonic/gin"
)

func drawRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	users := v1.Group("/users")
	{
		users.POST("", CreateUser)
		users.GET("/show", GetUser)
	}
	products := v1.Group("/products")
	{
		products.POST("", CreateProduct)
		products.GET("", FetchAllProducts)
		products.GET("/:id", FetchProduct)
		products.PUT("/:id", UpdateProduct)
		products.DELETE("/:id", DestroyProduct)
	}

	categories := v1.Group("/categories")
	{
		categories.POST("", CreateCategory)
		categories.GET("", FetchCategories)
		categories.PUT("/:id", UpdateCategory)
		categories.DELETE("/:id", DestroyCategory)
	}
	return router
}
