package dto

type CreateOrderItemInput struct {
	Quantity int     `json:"quantity" binding:"gte=0"`
	Amount   float64 `json:"price" binding:"gt=0"`
	//orderid?
	//order?
	//productid?
	//product?
}

type UpdateOrderItemInput struct {
	OldQuantity int `json:"old_quantity" binding:"required"`
	NewQuantity int `json:"new_quantity" binding:"required"`
}

//Implement order and product struct
