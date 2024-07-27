package routes

import (
	"sales-analysis/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/revenue", controllers.GetTotalRevenue)
}
