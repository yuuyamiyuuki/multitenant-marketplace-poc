package dto

type CheckoutInput struct {
	PaymentMethod string `json:"payment_method" binding:"required,oneof=pix payment_slip credit_card"`
}
