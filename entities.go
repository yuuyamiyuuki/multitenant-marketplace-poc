// Entities declaration

package main

import (
  "context"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Product struct {
	id datatypes.UUID
	name string
	price float
	stock int
	tags datatypes.JSON
	created_at date
	updated_at date
}

type Client struct{
	id datatypes.UUID
	document string
	phone string
	birthday date
	name string
	created_at date
	updated_at date
}

type Sale struct{
	id datatypes.UUID
	date date
	total_amount float
	details datatypes.JSON
	tentant_id datatypes.UUID
	client_id dataypes.UUID
	created_at date
	updated_at date
}

type Tenant struct {
	id datatypes.UUID
	name string
	settings datatypes.JSON
	birthday date
	apy_key datatypes.JSON
	created_at date
	updated_at date
}

type Tenant struct {
	id datatypes.UUID
	role enum
	username string
	password string
	tentant_id datatypes.UUID
	created_at date
	updated_at date
}

type SaleItems struct{
	product_id datatypes.UUID
	quatity int
	amount float
}
