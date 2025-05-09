package dto

import (
	"time"

	"github.com/adopabianko/dbo-service/internal/order/entity"
)

type OrderListRequest struct {
	Page   int
	Limit  int
	SortBy string
	Search string
}

type OrderListResponse struct {
	Meta any     `json:"meta"`
	Data []Order `json:"data"`
}

type Order struct {
	ID            string                 `json:"id"`
	Ref           string                 `json:"ref"`
	TotalQuantity float32                `json:"total_quantity"`
	TotalPrice    float32                `json:"total_price"`
	Customer      entity.OrderCustomer   `json:"customer"`
	Items         []entity.OrderItemList `json:"items"`
	CreatedAt     time.Time              `json:"created_at"`
	CreatedBy     string                 `json:"created_by"`
	UpdatedAt     *time.Time             `json:"updated_at"`
	UpdatedBy     *string                `json:"updated_by"`
}

type CreateOrderRequest struct {
	CustomerID string       `json:"customer_id" binding:"required"`
	OrderItems []orderItems `json:"items" binding:"required"`

	CreatedBy string `json:"created_by"`
}

type UpdateOrderRequest struct {
	CustomerID string       `json:"customer_id" binding:"required"`
	OrderItems []orderItems `json:"items" binding:"required"`

	ID        string
	UpdatedAt time.Time
	UpdatedBy string `json:"updated_by"`
}

type orderItems struct {
	ProductID string  `json:"product_id"`
	Quantity  float32 `json:"quantity"`
}

type Response struct {
	ID string `json:"id"`
}
