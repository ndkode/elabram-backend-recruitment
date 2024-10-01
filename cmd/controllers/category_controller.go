package controllers

import (
	"net/http"
	"strconv"

	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/services"
	"github.com/ndkode/elabram-backend-recruitment/cmd/utils"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	Service services.CategoryService
}

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetAllCategories(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
}

func NewCategoryController(service services.CategoryService) *categoryController {
	return &categoryController{Service: service}
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		reason := utils.HandleUnmarshalTypeError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": reason})
		return
	}

	// Validate category fields
	validationErrors := utils.ValidateStruct(category)
	if validationErrors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	if err := c.Service.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (c *categoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.Service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.Service.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}
