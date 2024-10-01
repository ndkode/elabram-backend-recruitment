package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPostProductsRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().CreateProduct(gomock.Any()).Return(nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.POST("/products", productController.CreateProduct)

	// Create a new request
	payload := models.Product{
		Name:          "product 1",
		Description:   "product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/products", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "product 1")
}

func TestPostProductsRouteBadRequest(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.POST("/products", productController.CreateProduct)

	// Create a new request
	payload := models.Product{
		Name:          "te",
		Description:   "product description 1",
		Price:         0,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/products", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "errors")
}

func TestGetProductsRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetAllProducts().Return([]models.Product{
		{
			Name:          "product 1",
			Description:   "product description 1",
			Price:         100,
			StockQuantity: 10,
			IsActive:      true,
			CategoryID:    1,
		},
		{
			Name:          "product 2",
			Description:   "product description 2",
			Price:         200,
			StockQuantity: 20,
			IsActive:      false,
			CategoryID:    2,
		},
	}, nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.GET("/products", productController.GetAllProducts)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "product 1")
}

func TestGetProductByIdRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetProductByID(gomock.Any()).Return(models.Product{
		Name:          "product 1",
		Description:   "product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}, nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.GET("/products/1", productController.GetProductByID)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "product 1")
}

func TestGetProductByIdRouteNotFound(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetProductByID(gomock.Any()).Return(models.Product{}, errors.New("not found"))

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.GET("/products/2", productController.GetProductByID)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/products/2", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "error")
}

func TestPutProductByIdRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetProductByID(gomock.Any()).Return(models.Product{
		Name:          "product 1",
		Description:   "product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}, nil)

	mockProductService.EXPECT().UpdateProduct(gomock.Any()).Return(models.Product{
		Name:          "product 1",
		Description:   "changed product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}, nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.PUT("/products/1", productController.UpdateProduct)

	// Create a new request
	payload := models.Product{
		Name:          "product 1",
		Description:   "changed product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/products/1", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "changed product description")
}

func TestPutProductByIdRouteBadRequest(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetProductByID(gomock.Any()).Return(models.Product{
		Name:          "product 1",
		Description:   "product description 1",
		Price:         100,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}, nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.PUT("/products/1", productController.UpdateProduct)

	// Create a new request
	payload := models.Product{
		Name:          "te",
		Description:   "changed product description 1",
		Price:         0,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/products/1", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "errors")
}

func TestPutProductByIdRouteNotFound(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().GetProductByID(gomock.Any()).Return(models.Product{}, errors.New("product not found"))

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.PUT("/products/2", productController.UpdateProduct)

	// Create a new request
	payload := models.Product{
		Name:          "te",
		Description:   "changed product description 1",
		Price:         0,
		StockQuantity: 10,
		IsActive:      true,
		CategoryID:    1,
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPut, "/products/2", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "error")
}

func TestDeleteProductByIdRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the ProductService
	mockProductService := mocks.NewMockProductService(ctrl)

	// Set up expectations
	mockProductService.EXPECT().DeleteProduct(gomock.Any()).Return(nil)

	// Set up the controller with the mocked service
	productController := controllers.NewProductController(mockProductService)
	r.DELETE("/products/1", productController.DeleteProduct)

	// Create a new request
	req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "deleted")
}
