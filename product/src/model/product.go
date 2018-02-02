package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Product
type Product struct {
	ID          int             `json:"id"`
	CategoryID  int             `json:"categoryId"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Stock       decimal.Decimal `json:"stock"`
	Price       decimal.Decimal `json:"price"`
	Version     int             `json:"version"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

// Products type
type Products []Product

// NewProduct function for initialise Product model
func NewProduct() *Product {
	now := time.Now()
	return &Product{
		Version:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
