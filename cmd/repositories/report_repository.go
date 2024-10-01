package repositories

import (
	"fmt"
	"sync"

	"github.com/ndkode/elabram-backend-recruitment/cmd/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reportRepository struct {
	DB *gorm.DB
}

type ReportRepository interface {
	GenerateProductReportWithGoroutines(ctx *gin.Context) (map[string]interface{}, error)
	GenerateProductReport(ctx *gin.Context) (map[string]interface{}, error)
}

func NewReportRepository(db *gorm.DB) *reportRepository {
	return &reportRepository{DB: db}
}

func (r *reportRepository) GenerateProductReportWithGoroutines(ctx *gin.Context) (map[string]interface{}, error) {
	fmt.Println("GenerateProductReportWithGoroutines")

	// Create channels to receive data from each goroutine
	totalProductsChan := make(chan int)
	totalStockChan := make(chan int)
	avgPriceChan := make(chan float64)
	productsChan := make(chan []models.Product)
	errChan := make(chan error, 1) // To catch errors
	var wg sync.WaitGroup
	wg.Add(4)

	// Query for total number of products, total stock, and average price
	go func() {
		defer wg.Done()
		var totalProducts int64
		// Apply filters
		db := applyFilters(ctx, r.DB)
		if err := db.Model(&models.Product{}).Count(&totalProducts).Error; err != nil {
			errChan <- err
			return
		}
		totalProductsChan <- int(totalProducts)
	}()

	go func() {
		defer wg.Done()
		var totalStock int64
		// Apply filters
		db := applyFilters(ctx, r.DB)
		if err := db.Model(&models.Product{}).Select("COALESCE(SUM(stock_quantity), 0)").Scan(&totalStock).Error; err != nil {
			errChan <- err
			return
		}
		totalStockChan <- int(totalStock)
	}()

	go func() {
		defer wg.Done()
		var avgPrice float64
		// Apply filters
		db := applyFilters(ctx, r.DB)
		if err := db.Model(&models.Product{}).Select("COALESCE(AVG(price), 0)").Scan(&avgPrice).Error; err != nil {
			errChan <- err
			return
		}
		avgPriceChan <- avgPrice
	}()

	// Get product details (with selected columns for efficiency)
	go func() {
		defer wg.Done()
		var products []models.Product

		// Apply pagination, sorting
		db := applyFilters(ctx, r.DB)
		db = applySorting(ctx, db)
		db, _, _ = applyPagination(ctx, db)
		if err := db.Preload("Category").Select("id, name, price, stock_quantity, category_id").Find(&products).Error; err != nil {
			errChan <- err
			return
		}
		productsChan <- products
	}()

	// Close channels once all goroutines finish
	go func() {
		wg.Wait()
		close(totalProductsChan)
		close(totalStockChan)
		close(avgPriceChan)
		close(productsChan)
		close(errChan) // Close error channel
	}()

	// Check for errors
	select {
	case err := <-errChan:
		return nil, err
	default:
		// No error, proceed with reading data from channels
	}

	// Read data from channels
	totalProducts := <-totalProductsChan
	totalStock := <-totalStockChan
	avgPrice := <-avgPriceChan
	products := <-productsChan

	return map[string]interface{}{
		"total_products": totalProducts,
		"total_stock":    totalStock,
		"avg_price":      avgPrice,
		"products":       products,
	}, nil
}

func (r *reportRepository) GenerateProductReport(ctx *gin.Context) (map[string]interface{}, error) {
	fmt.Println("GenerateProductReport")
	var (
		totalProducts int64
		totalStock    int64
		avgPrice      float64
		products      []models.Product
	)

	// Apply filters, sorting
	db := applyFilters(ctx, r.DB)
	db = applySorting(ctx, db)

	// Query for total number of products, total stock, and average price
	db.Model(&models.Product{}).Count(&totalProducts)
	db.Model(&models.Product{}).Select("COALESCE(SUM(stock_quantity), 0)").Scan(&totalStock)
	db.Model(&models.Product{}).Select("COALESCE(AVG(price), 0)").Scan(&avgPrice)

	// Apply pagination
	db, _, _ = applyPagination(ctx, db)

	// Get product details (with selected columns for efficiency)
	db.Preload("Category").Select("id, name, price, stock_quantity, category_id").Find(&products)

	return map[string]interface{}{
		"total_products": totalProducts,
		"total_stock":    totalStock,
		"avg_price":      avgPrice,
		"products":       products,
	}, nil
}

// Function to apply filters based on query parameters
func applyFilters(ctx *gin.Context, db *gorm.DB) *gorm.DB {
	// Filter by product name
	if name := ctx.Query("name"); name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	// Filter by category
	if categoryID := ctx.Query("category_id"); categoryID != "" {
		db = db.Where("category_id = ?", categoryID)
	}

	// Filter by price range
	if minPrice := ctx.Query("min_price"); minPrice != "" {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice := ctx.Query("max_price"); maxPrice != "" {
		db = db.Where("price <= ?", maxPrice)
	}

	// Filter by stock range
	if minStock := ctx.Query("min_stock"); minStock != "" {
		db = db.Where("stock_quantity >= ?", minStock)
	}
	if maxStock := ctx.Query("max_stock"); maxStock != "" {
		db = db.Where("stock_quantity <= ?", maxStock)
	}

	return db
}

// Function to apply sorting based on query parameters
func applySorting(ctx *gin.Context, db *gorm.DB) *gorm.DB {
	// Sorting by column name and order (ascending or descending)
	sortBy := ctx.DefaultQuery("sort_by", "name")
	sortOrder := ctx.DefaultQuery("sort_order", "asc")

	validSortColumns := map[string]bool{
		"name":           true,
		"category_id":    true,
		"price":          true,
		"stock_quantity": true,
	}

	if validSortColumns[sortBy] {
		db = db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
	}

	return db
}
