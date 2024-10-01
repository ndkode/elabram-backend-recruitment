package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"
)

func CategoryRoutes(router *gin.Engine, categoryController controllers.CategoryController) {
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("", categoryController.CreateCategory)
		categoryRoutes.GET("", categoryController.GetAllCategories)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
	}
}
