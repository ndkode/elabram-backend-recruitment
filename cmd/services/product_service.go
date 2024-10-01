package services

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"

	"github.com/gin-gonic/gin"
)

type productService struct {
	Repo repositories.ProductRepository
}

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetAllProductsWithPagination(ctx *gin.Context) (models.ProductsPageable, error)
	GetProductByID(id uint) (models.Product, error)
	UpdateProduct(product *models.Product) (models.Product, error)
	DeleteProduct(id uint) error
}

func NewProductService(repo repositories.ProductRepository) *productService {
	return &productService{Repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *productService) GetAllProductsWithPagination(ctx *gin.Context) (models.ProductsPageable, error) {
	return s.Repo.GetAllProductsWithPagination(ctx)
}

func (s *productService) GetProductByID(id uint) (models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *productService) UpdateProduct(product *models.Product) (models.Product, error) {
	return s.Repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
