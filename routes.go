package main

import (
	. "coconut/controller"

	"github.com/gin-gonic/gin"
)

func drawRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/products")
	{
		v1.POST("/", CreateProduct)
		v1.GET("/", FetchAllProducts)
		v1.GET("/:id", FetchProduct)
		v1.PUT("/:id", UpdateProduct)
		v1.DELETE("/:id", DestroyProduct)
	}
	return router
}
