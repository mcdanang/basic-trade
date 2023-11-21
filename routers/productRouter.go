package routers

import (
	"basic-trade/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	productRouter := router.Group("/products")
	{
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetAllProduct)
	}

	// router.PUT("/orders/:orderID", controllers.UpdateOrder)

	// router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
