package service

import (
	"context"
	"math"

	"github.com/adopabianko/dbo-service/internal/customer/dto"
	"github.com/adopabianko/dbo-service/internal/customer/entity"
	"github.com/adopabianko/dbo-service/pkg/http/response"
)

func (s *service) FindAll(ctx context.Context, params dto.CustomerListRequest) (dto.CustomerListResponse, error) {
	var result dto.CustomerListResponse

	customers, err := s.repository.FindAll(ctx, params)
	if err != nil {
		return result, err
	}

	var totalItems int
	if len(customers) > 0 {
		totalItems = customers[0].Total
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	result.Data = customers
	result.Meta = response.BuildMeta(
		params.Limit,
		totalItems,
		params.Page,
		totalPages,
		params.SortBy,
	)

	return result, nil
}

func (s *service) FindByID(ctx context.Context, id string) (entity.Customer, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Create(ctx context.Context, params dto.CreateCustomerRequest) (dto.Response, error) {
	var customer entity.Customer

	customer.Name = params.Name
	customer.Phone = params.Phone
	customer.Email = params.Email
	customer.Gender = params.Gender
	customer.Address = params.Address
	customer.CreatedBy = params.CreatedBy

	create, err := s.repository.Create(ctx, customer)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.Response{ID: create.ID}, nil
}

func (s *service) Update(ctx context.Context, params dto.UpdateCustomerRequest) (dto.Response, error) {
	customer, err := s.repository.FindByID(ctx, params.ID)
	if err != nil {
		return dto.Response{}, err
	}

	customer.Name = params.Name
	customer.Phone = params.Phone
	customer.Email = params.Email
	customer.Gender = params.Gender
	customer.Address = params.Address
	customer.UpdatedBy = &params.UpdatedBy

	update, err := s.repository.Update(ctx, customer)
	if err != nil {
		return dto.Response{}, err
	}
	return dto.Response{ID: update.ID}, nil
}

func (s *service) Delete(ctx context.Context, email, id string) error {
	return s.repository.Delete(ctx, email, id)
}
