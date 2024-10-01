package services

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"
)

type categoryService struct {
	Repo repositories.CategoryRepository
}

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

func NewCategoryService(repo repositories.CategoryRepository) *categoryService {
	return &categoryService{Repo: repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.Repo.CreateCategory(category)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.Repo.GetAllCategories()
}

func (s *categoryService) GetCategoryByID(id uint) (models.Category, error) {
	return s.Repo.GetCategoryByID(id)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.Repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.Repo.DeleteCategory(id)
}
