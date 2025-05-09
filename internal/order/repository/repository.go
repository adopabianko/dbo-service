package repository

import (
	"context"

	"github.com/adopabianko/dbo-service/internal/order/dto"
	"github.com/adopabianko/dbo-service/internal/order/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	IOrderRepository interface {
		// Order repository interface
		FindAll(ctx context.Context, params dto.OrderListRequest) ([]*entity.OrderList, error)
		FindByID(ctx context.Context, ID string) (*entity.OrderList, error)
		Create(ctx context.Context, order entity.Order) (string, error)
		Update(ctx context.Context, order entity.Order) (string, error)
		Delete(ctx context.Context, email, ID string) error

		// OrderItem repository interface
		FindOrderItemsByOrderId(ctx context.Context, orderID string) ([]entity.OrderItemList, error)
		CreateOrderItem(ctx context.Context, orderItem entity.OrderItem) (string, error)
		DeleteOrderItem(ctx context.Context, orderID string) error
		FindProductByID(ctx context.Context, ID string) (*entity.Product, error)
	}

	repository struct {
		db *pgxpool.Pool
	}
)

func NewRepository(db *pgxpool.Pool) IOrderRepository {
	return &repository{
		db: db,
	}
}
