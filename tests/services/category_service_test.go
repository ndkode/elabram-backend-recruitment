package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/services"
	"github.com/ndkode/elabram-backend-recruitment/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockCategoryRepository(ctrl)
	service := services.NewCategoryService(mockRepository)

	var categories []models.Category
	mockRepository.EXPECT().GetAllCategories().Return(categories, nil).Times(1)

	result, err := service.GetAllCategories()

	assert.Nil(t, err)
	assert.Equal(t, categories, result)
}

func TestGetCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockCategoryRepository(ctrl)
	service := services.NewCategoryService(mockRepository)

	category := models.Category{ID: 1}
	mockRepository.EXPECT().GetCategoryByID(uint(1)).Return(category, nil).Times(1)

	result, err := service.GetCategoryByID(uint(1))

	assert.Nil(t, err)
	assert.Equal(t, category, result)
}

func TestCreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockCategoryRepository(ctrl)
	service := services.NewCategoryService(mockRepository)

	category := models.Category{}
	mockRepository.EXPECT().CreateCategory(&category).Return(nil)

	err := service.CreateCategory(&category)

	assert.Nil(t, err)
}

func TestUpdateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockCategoryRepository(ctrl)
	service := services.NewCategoryService(mockRepository)

	category := models.Category{ID: 1}
	mockRepository.EXPECT().UpdateCategory(&category).Return(nil)

	err := service.UpdateCategory(&category)

	assert.Nil(t, err)
}

func TestDeleteCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockCategoryRepository(ctrl)
	service := services.NewCategoryService(mockRepository)

	mockRepository.EXPECT().DeleteCategory(uint(1)).Return(nil).Times(1)

	err := service.DeleteCategory(uint(1))

	assert.Nil(t, err)
}
