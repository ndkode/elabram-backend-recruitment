package controllers

import (
	"net/http"
	"strconv"

	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/services"
	"github.com/ndkode/elabram-backend-recruitment/cmd/utils"

	"github.com/gin-gonic/gin"
)

type productController struct {
	Service services.ProductService
}

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
	GetAllProductsWithPagination(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

func NewProductController(service services.ProductService) *productController {
	return &productController{Service: service}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		reason := utils.HandleUnmarshalTypeError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": reason})
		return
	}

	// Validate product fields
	validationErrors := utils.ValidateStruct(product)
	if validationErrors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	if err := c.Service.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *productController) GetAllProducts(ctx *gin.Context) {
	products, err := c.Service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *productController) GetAllProductsWithPagination(ctx *gin.Context) {
	productsWithPagination, err := c.Service.GetAllProductsWithPagination(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, productsWithPagination)
}

func (c *productController) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.Service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	var updateProduct models.Product
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&updateProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.Service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// update only specific product fields, to prevent accidental changes to other fields
	if updateProduct.Name != "" {
		product.Name = updateProduct.Name
	}
	if updateProduct.Description != "" {
		product.Description = updateProduct.Description
	}
	if updateProduct.Price != product.Price {
		product.Price = updateProduct.Price
	}
	if updateProduct.CategoryID != product.CategoryID {
		product.CategoryID = updateProduct.CategoryID
	}
	if updateProduct.StockQuantity != product.StockQuantity {
		product.StockQuantity = updateProduct.StockQuantity
	}
	if updateProduct.IsActive != product.IsActive {
		product.IsActive = updateProduct.IsActive
	}

	// Validate product fields
	validationErrors := utils.ValidateStruct(product)
	if validationErrors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	updatedProduct, err := c.Service.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Service.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
