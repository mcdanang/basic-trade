package routers

import (
	"basic-trade/controllers"
	"basic-trade/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.AdminRegister)
		authRouter.POST("/login", controllers.AdminLogin)
	}

	productRouter := router.Group("/products")
	{

		productRouter.GET("/variants", controllers.GetAllVariant)
		productRouter.GET("/variants/:variantUUID", controllers.GetVariantByUUID)

		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productUUID", controllers.GetProductByUUID)

		productRouter.Use(middlewares.Authentication())

		productRouter.POST("/variants", controllers.CreateVariant)
		productRouter.PUT("/variants/:variantUUID", controllers.UpdateVariant)
		productRouter.DELETE("/variants/:variantUUID", controllers.DeleteVariant)

		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productUUID", controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", controllers.DeleteProduct)
	}

	return router
}
