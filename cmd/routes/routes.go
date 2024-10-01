package routes

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/configs"
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"
	"github.com/ndkode/elabram-backend-recruitment/cmd/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Initialize Repository, Service, and Controller
	productRepo := repositories.NewProductRepository(configs.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)
	ProductRoutes(r, productController)

	categoryRepo := repositories.NewCategoryRepository(configs.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)
	CategoryRoutes(r, categoryController)
}
