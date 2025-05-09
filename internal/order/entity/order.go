package entity

import "time"

type Order struct {
	ID            string     `json:"id" db:"id"`
	Ref           string     `json:"name" db:"ref"`
	CustomerID    string     `json:"customer_id" db:"customer_id"`
	TotalQuantity float32    `json:"total_quantity" db:"total_quantity"`
	TotalPrice    float32    `json:"total_price" db:"total_price"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	CreatedBy     string     `json:"created_by" db:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy     *string    `json:"updated_by" db:"updated_by"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	DeletedBy     *string    `json:"deleted_by" db:"deleted_by"`

	Total int `json:"-" db:"total"`
}

type OrderItem struct {
	ID        string  `json:"id" db:"id"`
	OrderID   string  `json:"order_id" db:"order_id"`
	ProductID string  `json:"product_id" db:"product_id"`
	Quantity  float32 `json:"quantity" db:"quantity"`
	Subtotal  float32 `json:"subtotal" db:"subtotal"`
}

type OrderList struct {
	ID            string        `json:"id" db:"id"`
	Ref           string        `json:"name" db:"ref"`
	Customer      OrderCustomer `json:"customer"`
	TotalQuantity float32       `json:"total_quantity" db:"total_quantity"`
	TotalPrice    float32       `json:"total_price" db:"total_price"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	CreatedBy     string        `json:"created_by" db:"created_by"`
	UpdatedAt     *time.Time    `json:"updated_at" db:"updated_at"`
	UpdatedBy     *string       `json:"updated_by" db:"updated_by"`

	Total int `json:"-" db:"total"`
}

type OrderCustomer struct {
	CustomerID    string `json:"customer_id" db:"customer_id"`
	CustomerName  string `json:"customer_name" db:"customer_name"`
	CustomerEmail string `json:"customer_email" db:"customer_email"`
	CustomerPhone string `json:"customer_phone" db:"customer_phone"`
}

type OrderItemList struct {
	ID       string           `json:"id" db:"id"`
	OrderID  string           `json:"-" db:"order_id"`
	Product  OrderItemProduct `json:"product"`
	Quantity float32          `json:"quantity"`
	Subtotal float32          `json:"subtotal"`
}

type OrderItemProduct struct {
	ProductID   string `json:"product_id"`
	ProductSku  string `json:"product_sku"`
	ProductName string `json:"product_name"`
}

type Product struct {
	ID    string  `json:"id" db:"id"`
	Sku   string  `json:"sku" db:"sku"`
	Name  string  `json:"name" db:"name"`
	Price float32 `json:"price" db:"price"`
}
