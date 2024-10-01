package models

import (
	"time"
)

type Product struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name" validate:"required,min=3,max=100"`
	Description   string    `json:"description"`
	Price         float64   `json:"price" validate:"required,gt=0"`
	CategoryID    uint      `json:"category_id" validate:"required"`
	Category      *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	StockQuantity int       `json:"stock_quantity" validate:"gt=0"`
	IsActive      bool      `json:"is_active" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductsPageable struct {
	Products   []Product `json:"products"`
	Page       int       `json:"page"`
	TotalItems int64     `json:"total_items"`
	TotalPages int       `json:"total_pages"`
}
