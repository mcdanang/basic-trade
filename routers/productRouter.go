package routers

import (
	"basic-trade/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	variantRouter := router.Group("/products/variant")
	{
		// variantRouter.POST("/", controllers.CreateProduct)
		variantRouter.GET("/", controllers.GetAllVariant)
		// variantRouter.GET("/:productUUID", controllers.GetProductByUUID)
		// variantRouter.PUT("/:productUUID", controllers.UpdateProduct)
		// variantRouter.DELETE("/:productUUID", controllers.DeleteProduct)
	}

	productRouter := router.Group("/products")
	{
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productUUID", controllers.GetProductByUUID)
		productRouter.PUT("/:productUUID", controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", controllers.DeleteProduct)
	}

	return router
}
