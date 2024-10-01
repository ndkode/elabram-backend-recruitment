package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCategoryRepository(t *testing.T) {
	// Load the database configuration
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "rootpassword")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "elabram")
	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Category{})

	// Create a new repository
	repo := repositories.NewCategoryRepository(db)

	// Test Create
	category := models.Category{Name: "Test Category", Description: "Test Category Description"}
	err = repo.CreateCategory(&category)
	assert.NoError(t, err)

	// Test GetAll
	categories, err := repo.GetAllCategories()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(categories), 0)
	assert.Equal(t, category.Name, categories[len(categories)-1].Name)

	// Test GetByID
	categoryByID, err := repo.GetCategoryByID(categories[len(categories)-1].ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, categoryByID.Name)

	// Test Update
	category.Name = "Updated Test Category"
	err = repo.UpdateCategory(&category)
	assert.NoError(t, err)

	// Test Delete
	err = repo.DeleteCategory(categories[len(categories)-1].ID)
	assert.NoError(t, err)
}
