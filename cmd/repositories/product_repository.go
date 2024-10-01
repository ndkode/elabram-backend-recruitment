package repositories

import (
	"math"
	"strconv"

	"github.com/ndkode/elabram-backend-recruitment/cmd/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetAllProductsWithPagination(ctx *gin.Context) (models.ProductsPageable, error)
	GetProductByID(id uint) (models.Product, error)
	UpdateProduct(product *models.Product) (models.Product, error)
	DeleteProduct(id uint) error
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) GetAllProductsWithPagination(ctx *gin.Context) (models.ProductsPageable, error) {
	db, page, pageSize := applyPagination(ctx, r.DB)

	productsPageable := models.ProductsPageable{}

	err := db.Preload("Category").Find(&productsPageable.Products).Error
	r.DB.Model(&models.Product{}).Count(&productsPageable.TotalItems)
	productsPageable.TotalPages = int(math.Ceil(float64(productsPageable.TotalItems) / float64(pageSize)))
	productsPageable.Page = page

	return productsPageable, err
}

func (r *productRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Category").First(&product, id).Error
	return product, err
}

func (r *productRepository) UpdateProduct(product *models.Product) (models.Product, error) {
	updatedProduct := models.Product{}
	product.Category = nil
	err := r.DB.Where("id = ?", product.ID).Updates(&product).Preload("Category").First(&updatedProduct).Error
	return updatedProduct, err
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}

// Function to apply pagination based on query parameters
func applyPagination(ctx *gin.Context, db *gorm.DB) (*gorm.DB, int, int) {
	// Default page and page size
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	offset := (page - 1) * pageSize
	return db.Offset(offset).Limit(pageSize), page, pageSize
}
