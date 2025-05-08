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

func (s *service) Create(ctx context.Context, params dto.CreateCustomerRequest) (entity.Customer, error) {
	var opt entity.Customer

	opt.Name = params.Name
	opt.Phone = params.Phone
	opt.Email = params.Email
	opt.Gender = params.Gender
	opt.Address = params.Address
	opt.CreatedBy = params.CreatedBy

	return s.repository.Create(ctx, opt)
}

func (s *service) Update(ctx context.Context, params dto.UpdateCustomerRequest) (entity.Customer, error) {
	opt, err := s.repository.FindByID(ctx, params.ID)
	if err != nil {
		return entity.Customer{}, err
	}

	opt.Name = params.Name
	opt.Phone = params.Phone
	opt.Email = params.Email
	opt.Gender = params.Gender
	opt.Address = params.Address
	opt.UpdatedAt = &params.UpdatedAt
	opt.UpdatedBy = &params.UpdatedBy

	return s.repository.Update(ctx, opt)
}

func (s *service) Delete(ctx context.Context, id string) error {
	_, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repository.Delete(ctx, id)
}
