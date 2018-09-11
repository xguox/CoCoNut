package main

import (
	. "coconut/controller"
	"coconut/util"

	"github.com/gin-gonic/gin"
)

func drawRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/")

	shop := v1.Group("/shop")
	shopRoutesRegister(shop)

	admin := v1.Group("/admin")
	adminRoutesRegister(admin)

	return router
}

func adminRoutesRegister(admin *gin.RouterGroup) {
	admin.Use(util.AuthMiddleware())

	{
		admin.POST("login", UserLogin)
	}

	users := admin.Group("/users")
	{
		users.POST("", CreateUser)
		users.GET("/show", GetUser)
	}

	products := admin.Group("/products")
	{
		products.POST("", CreateProduct)
		products.GET("", FetchAllProducts)
		products.GET("/:id", FetchProduct)
		products.PUT("/:id", UpdateProduct)
		products.DELETE("/:id", DestroyProduct)
	}

	categories := admin.Group("/categories")
	{
		categories.POST("", CreateCategory)
		categories.GET("", FetchCategories)
		categories.PUT("/:id", UpdateCategory)
		categories.DELETE("/:id", DestroyCategory)
	}
}

func shopRoutesRegister(shop *gin.RouterGroup) {

}
