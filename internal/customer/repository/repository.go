package repository

import (
	"context"

	"github.com/adopabianko/dbo-service/internal/customer/dto"
	"github.com/adopabianko/dbo-service/internal/customer/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	ICustomerRepository interface {
		FindAll(ctx context.Context, params dto.CustomerListRequest) ([]entity.Customer, error)
		FindByID(ctx context.Context, ID string) (entity.Customer, error)
		Create(ctx context.Context, customer entity.Customer) (entity.Customer, error)
		Update(ctx context.Context, customer entity.Customer) (entity.Customer, error)
		Delete(ctx context.Context, ID string) error
	}

	repository struct {
		db *pgxpool.Pool
	}
)

func NewRepository(db *pgxpool.Pool) ICustomerRepository {
	return &repository{
		db: db,
	}
}
