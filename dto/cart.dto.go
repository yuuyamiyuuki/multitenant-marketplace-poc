package dto

import "github.com/google/uuid"

type CartItemInput struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type UpsertCartInput struct {
	Items []CartItemInput `json:"items"`
}

type CartItemResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
}

type CartResponse struct {
	ID         uuid.UUID          `json:"id"`
	TotalItems int                `json:"total_items"`
	TotalPrice float64            `json:"total_price"`
	Items      []CartItemResponse `json:"items"`
}
