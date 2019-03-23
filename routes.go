package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	. "github.com/xguox/coconut/controller"
	_ "github.com/xguox/coconut/docs"
	"github.com/xguox/coconut/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func drawRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	v1 := router.Group("/api/v1/")

	shop := v1.Group("/shop")
	shopRoutesRegister(shop)

	admin := v1.Group("/admin")
	adminRoutesRegister(admin)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func adminRoutesRegister(admin *gin.RouterGroup) {
	admin.POST("login", UserLogin)
	admin.POST("/users", CreateUser)
	admin.Use(middleware.AuthMiddleware())

	users := admin.Group("/users")
	{
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

		products.GET("/:id/variants", ProductVariants)
		products.GET("/:id/variants/:variant_id", ProductVariant)
		products.PUT("/:id/variants/:variant_id", UpdateVariant)
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
