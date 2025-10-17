package models

import "time"

// Product representa un registro en la table 'products
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProducRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description *string `json:"description"`
}
