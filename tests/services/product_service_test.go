package services_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ndkode/elabram-backend-recruitment/cmd/models"
	"github.com/ndkode/elabram-backend-recruitment/cmd/services"
	"github.com/ndkode/elabram-backend-recruitment/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	product := models.Product{}
	mockRepository.EXPECT().CreateProduct(&product).Return(nil)

	err := service.CreateProduct(&product)

	assert.Nil(t, err)
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	product := models.Product{ID: 1}
	mockRepository.EXPECT().UpdateProduct(&product).Return(product, nil).Times(1)

	result, err := service.UpdateProduct(&product)

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	mockRepository.EXPECT().DeleteProduct(uint(1)).Return(nil).Times(1)

	err := service.DeleteProduct(uint(1))

	assert.Nil(t, err)
}

func TestGetAllProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	products := []models.Product{}
	mockRepository.EXPECT().GetAllProducts().Return(products, nil).Times(1)

	result, err := service.GetAllProducts()

	assert.Nil(t, err)
	assert.Equal(t, products, result)
}

func TestGetAllProductsWithPagination(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	productsPageable := models.ProductsPageable{}
	mockRepository.EXPECT().GetAllProductsWithPagination(gomock.Any()).Return(productsPageable, nil)

	result, err := service.GetAllProductsWithPagination(&gin.Context{})

	assert.Nil(t, err)
	assert.Equal(t, productsPageable, result)
}

func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockProductRepository(ctrl)
	service := services.NewProductService(mockRepository)

	product := models.Product{}
	mockRepository.EXPECT().GetProductByID(uint(1)).Return(product, nil).Times(1)

	result, err := service.GetProductByID(uint(1))

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
