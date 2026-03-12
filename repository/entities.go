package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Name      string
	Settings  datatypes.JSON `json:"settings" swaggertype:"object"`
	APIKey    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []User
	Orders    []Order
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Role      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	TenantID  uuid.UUID
	Tenant    Tenant
}

type Client struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Document  string
	Phone     string
	Birthday  time.Time
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []Order
	UserId    uuid.UUID
	User      User
}

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Name      string
	Price     float64
	Stock     int
	Tags      datatypes.JSON `json:"tags" swaggertype:"string"`
	TenantID  uuid.UUID
	Tenant    Tenant
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CartID    uuid.UUID
	Cart      Cart
	ProductID uuid.UUID
	Product   Product
	Quantity  int
}

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	TenantID  uuid.UUID
	Tenant    Tenant
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []CartItem `gorm:"foreignKey:CartID"`
	UserID    uuid.UUID
	User      User
}

type Order struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Date          time.Time
	TotalAmount   float64
	PaymentMethod string
	Status        string
	Details       datatypes.JSON `json:"details" swaggertype:"object"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TenantID      uuid.UUID
	Tenant        Tenant
	ClientID      uuid.UUID
	Client        Client
	Items         []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	Quantity  int
	Amount    float64
	OrderID   uuid.UUID
	Order     Order
	ProductID uuid.UUID
	Product   Product
}
