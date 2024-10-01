package routes

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func ReportRoutes(router *gin.Engine, reportController controllers.ReportController) {
	categoryRoutes := router.Group("/reports")
	{
		categoryRoutes.GET("/products", reportController.GetProductReport)
	}
}
