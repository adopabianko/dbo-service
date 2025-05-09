package dto

import (
	"time"

	"github.com/adopabianko/dbo-service/internal/customer/entity"
)

type CustomerListRequest struct {
	Page   int
	Limit  int
	SortBy string
	Search string
}

type CustomerListResponse struct {
	Meta any               `json:"meta"`
	Data []entity.Customer `json:"data"`
}

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Address string `json:"address" binding:"required"`

	CreatedBy string `json:"created_by"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Address string `json:"address" binding:"required"`

	ID        string
	UpdatedAt time.Time
	UpdatedBy string `json:"updated_by"`
}

type Response struct {
	ID string `json:"id"`
}
