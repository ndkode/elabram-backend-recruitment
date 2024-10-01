package repositories

import (
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := r.DB.First(&category, id).Error
	return category, err
}

func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}
