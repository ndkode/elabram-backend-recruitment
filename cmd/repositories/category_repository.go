package repositories

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	return category, err
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}
