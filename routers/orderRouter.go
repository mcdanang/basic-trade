package routers

import (
	"assignment-project-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)

	router.GET("/orders", controllers.GetAllOrder)

	router.PUT("/orders/:orderID", controllers.UpdateOrder)

	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
