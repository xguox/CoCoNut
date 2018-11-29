package main

import (
	. "coconut/controller"
	"coconut/middleware"

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
	admin.Use(middleware.AuthMiddleware())

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
		products.POST("/:id/tagging", TaggingProduct)

		products.GET("/:id/options", FetchOptions)
		products.POST("/:id/init-build-options", InitBuildOptions)
		products.POST("/:id/create-option", CreateOption)
		products.POST("/:id/options/:option_id/add-single-val", AddSingleValue)
		products.DELETE("/:id/options/:option_id/del-val", DeleteSingleValue)
		products.DELETE("/:id/options/:option_id", DeleteOption)
		products.POST("/:id/reorder-options", ReorderOptions)

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
