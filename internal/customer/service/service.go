package service

import (
	"context"

	"github.com/adopabianko/dbo-service/internal/customer/dto"
	"github.com/adopabianko/dbo-service/internal/customer/entity"
	"github.com/adopabianko/dbo-service/internal/customer/repository"
)

type (
	ICustomerService interface {
		FindAll(ctx context.Context, params dto.CustomerListRequest) (dto.CustomerListResponse, error)
		FindByID(ctx context.Context, id string) (entity.Customer, error)
		Create(ctx context.Context, params dto.CreateCustomerRequest) (entity.Customer, error)
		Update(ctx context.Context, params dto.UpdateCustomerRequest) (entity.Customer, error)
		Delete(ctx context.Context, id string) error
	}

	service struct {
		repository repository.ICustomerRepository
	}
)

func NewService(customerRepo repository.ICustomerRepository) ICustomerService {
	return &service{
		repository: customerRepo,
	}
}
