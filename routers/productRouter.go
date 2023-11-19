package routers

import (
	"basic-trade/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// router.POST("/orders", controllers.CreateOrder)

	router.GET("/products", controllers.GetAllProduct)

	// router.PUT("/orders/:orderID", controllers.UpdateOrder)

	// router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
