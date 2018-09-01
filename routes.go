package main

import (
	. "coconut/controller"

	"github.com/gin-gonic/gin"
)

func drawRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/")
	products := v1.Group("/products")
	{
		products.POST("/", CreateProduct)
		products.GET("/", FetchAllProducts)
		products.GET("/:id", FetchProduct)
		products.PUT("/:id", UpdateProduct)
		products.DELETE("/:id", DestroyProduct)
	}
	return router
}
