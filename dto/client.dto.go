package dto

import (
	"time"
)

type CreateClientInput struct {
	Username string    `json:"username" binding:"required"`
	Document string    `json:"document" binding:"required,required,min=11, max=11"`
	Name     string    `json:"name" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Password string    `json:"password" binding:"required,min=6"`
}
