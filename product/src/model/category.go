package model

import "time"

// Category
type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Categories type
type Categories []Category

// NewCategory function for initialise Product model
func NewCategory() *Category {
	now := time.Now()
	return &Category{
		Version:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
