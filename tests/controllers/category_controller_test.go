package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ndkode/elabram-backend-recruitment/cmd/controllers"
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostCategoryRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the CategoryService
	mockCategoryService := mocks.NewMockCategoryService(ctrl)

	// Set up expectations
	mockCategoryService.EXPECT().CreateCategory(gomock.Any()).Return(nil)

	// Set up the controller with the mocked service
	categoryController := controllers.NewCategoryController(mockCategoryService)
	r.POST("/categories", categoryController.CreateCategory)

	// Create a new request
	payload := models.Category{
		Name:        "category 1",
		Description: "category product description 1",
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/categories", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "category 1")
	assert.Contains(t, recorder.Body.String(), "category product description 1")
}

func TestPostCategoryRouteBadRequest(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the CategoryService
	mockCategoryService := mocks.NewMockCategoryService(ctrl)

	// Set up the controller with the mocked service
	categoryController := controllers.NewCategoryController(mockCategoryService)
	r.POST("/categories", categoryController.CreateCategory)

	// Create a new request
	payload := models.Category{
		Name:        "",
		Description: "category product description 1",
	}
	payloadJson, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/categories", strings.NewReader(string(payloadJson)))

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "errors")
}

func TestGetAllCategoriesRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the CategoryService
	mockCategoryService := mocks.NewMockCategoryService(ctrl)

	// Set up expectations
	mockCategoryService.EXPECT().GetAllCategories().Return([]models.Category{
		{
			ID:          1,
			Name:        "category 1",
			Description: "category product description 1",
		},
		{
			ID:          2,
			Name:        "category 2",
			Description: "category product description 2",
		},
	}, nil)

	// Set up the controller with the mocked service
	categoryController := controllers.NewCategoryController(mockCategoryService)
	r.GET("/categories", categoryController.GetAllCategories)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/categories", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "category 1")
}

func TestGetCategoryByIdRoute(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the CategoryService
	mockCategoryService := mocks.NewMockCategoryService(ctrl)

	// Set up expectations
	mockCategoryService.EXPECT().GetCategoryByID(gomock.Any()).Return(models.Category{
		ID:          1,
		Name:        "category 1",
		Description: "category product description 1",
	}, nil)

	// Set up the controller with the mocked service
	categoryController := controllers.NewCategoryController(mockCategoryService)
	r.GET("/categories/1", categoryController.GetCategoryByID)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/categories/1", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "category 1")
}

func TestGetCategoryByIdRouteNotFound(t *testing.T) {
	r := gin.Default()
	recorder := httptest.NewRecorder()

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the CategoryService
	mockCategoryService := mocks.NewMockCategoryService(ctrl)

	// Set up expectations
	mockCategoryService.EXPECT().GetCategoryByID(gomock.Any()).Return(models.Category{}, errors.New("not found"))

	// Set up the controller with the mocked service
	categoryController := controllers.NewCategoryController(mockCategoryService)
	r.GET("/categories/2", categoryController.GetCategoryByID)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/categories/2", nil)

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "error")
}
