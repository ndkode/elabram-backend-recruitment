package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestProductRepository(t *testing.T) {
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
	db.AutoMigrate(&models.Product{})

	// Create a new repository
	repo := repositories.NewProductRepository(db)

	// Test Create
	product := models.Product{
		Name:          "Test Product",
		Description:   "Test Product Description",
		Price:         1000,
		StockQuantity: 10,
		CategoryID:    1,
		IsActive:      true,
	}
	err = repo.CreateProduct(&product)
	assert.NoError(t, err)

	// Test GetAll
	products, err := repo.GetAllProducts()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(products), 0)
	assert.Equal(t, product.Name, products[len(products)-1].Name)

	// Test GetAllWithPagination
	productsPageable, err := repo.GetAllProductsWithPagination(&gin.Context{})
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, int(productsPageable.TotalItems), 0)

	// Test GetByID
	productByID, err := repo.GetProductByID(products[len(products)-1].ID)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productByID.Name)

	// Test Update
	product.Name = "Updated Test Product"
	result, err := repo.UpdateProduct(&product)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, result.Name)

	// Test Delete
	err = repo.DeleteProduct(products[len(products)-1].ID)
	assert.NoError(t, err)
}
