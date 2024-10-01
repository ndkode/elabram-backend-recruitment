package services

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *ProductService) GetAllProductsWithPagination(ctx *gin.Context) (models.ProductsPageable, error) {
	return s.Repo.GetAllProductsWithPagination(ctx)
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) (models.Product, error) {
	return s.Repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.DeleteProduct(id)
}
