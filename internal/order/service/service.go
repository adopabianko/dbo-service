package service

import (
	"context"

	customerRepository "github.com/adopabianko/dbo-service/internal/customer/repository"
	"github.com/adopabianko/dbo-service/internal/order/dto"
	"github.com/adopabianko/dbo-service/internal/order/repository"
)

type (
	IOrderService interface {
		FindAll(ctx context.Context, params dto.OrderListRequest) (dto.OrderListResponse, error)
		FindByID(ctx context.Context, id string) (dto.Order, error)
		Create(ctx context.Context, params dto.CreateOrderRequest) (dto.Response, error)
		Update(ctx context.Context, params dto.UpdateOrderRequest) (dto.Response, error)
		Delete(ctx context.Context, email, id string) error
	}

	service struct {
		repository         repository.IOrderRepository
		customerRepository customerRepository.ICustomerRepository
	}
)

func NewService(
	orderRepo repository.IOrderRepository,
	customerRepo customerRepository.ICustomerRepository,
) IOrderService {
	return &service{
		repository:         orderRepo,
		customerRepository: customerRepo,
	}
}
