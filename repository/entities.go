package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string
	Settings  datatypes.JSON
	Birthday  time.Time
	APIKey    datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User
	Sales     []Sale
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Role      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	TenantID  uuid.UUID
	Tenant    Tenant
}

type Client struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Document  string
	Phone     string
	Birthday  time.Time
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Sales     []Sale
}

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string
	Price     float64
	Stock     int
	Tags      datatypes.JSON
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Sale struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Date        time.Time
	TotalAmount float64
	Details     datatypes.JSON
	CreatedAt   time.Time
	UpdatedAt   time.Time
	TenantID    uuid.UUID
	Tenant      Tenant
	ClientID    uuid.UUID
	Client      Client
	Items       []SaleItem `gorm:"foreignKey:SaleID"`
}

type SaleItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Quantity  int
	Amount    float64
	SaleID    uuid.UUID
	Sale      Sale
	ProductID uuid.UUID
	Product   Product
}
