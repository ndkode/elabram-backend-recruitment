package routes

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, productController controllers.ProductController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("", productController.CreateProduct)
		productRoutes.GET("", productController.GetAllProductsWithPagination)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}
}
