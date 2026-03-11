package dto

type CreateSaleItemInput struct {
	Quantity int     `json:"quantity" binding:"gte=0"`
	Amount   float64 `json:"price" binding:"gt=0"`
	//saleid?
	//sale?
	//productid?
	//product?
}

type UpdateSaleItemInput struct {
	OldQuantity int `json:"old_quantity" binding:"required"`
	NewQuantity int `json:"new_quantity" binding:"required"`
}

//Implement sale and product struct
