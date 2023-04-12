package routes

import (
	"gin_server/controllers"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.POST("/portfolio/v1/services/", controllers.GetPortfolioServices())
}