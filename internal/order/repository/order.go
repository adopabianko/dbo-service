package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/adopabianko/dbo-service/internal/order/dto"
	"github.com/adopabianko/dbo-service/internal/order/entity"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
)

func (r *repository) FindAll(ctx context.Context, params dto.OrderListRequest) ([]*entity.OrderList, error) {
	query := `
        SELECT
 			COUNT(*) OVER () as total,
			"order".id,
            "order".ref,
			customer.id as customer_id,
			customer.name as customer_name,
			customer.email as customer_email,
			customer.phone as customer_phone,
			"order".total_quantity,
			"order".total_price,
            "order".created_at,
            "order".created_by,
            "order".updated_at,
            "order".updated_by
        FROM "order"
		JOIN customer ON "order".customer_id = customer.id
        WHERE "order".deleted_at IS NULL`

	if params.Search != "" {
		query += fmt.Sprintf(` 
		AND ("order".id::text ILIKE '%s' 
		OR "order".ref ILIKE '%s' 
		OR customer.name ILIKE '%s' 
		OR customer.email ILIKE '%s')`,
			"%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	query += fmt.Sprintf(` ORDER BY %s LIMIT %d OFFSET %d`, params.SortBy, params.Limit, (params.Page-1)*params.Limit)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all orders")
	}
	defer rows.Close()

	var orders []*entity.OrderList
	for rows.Next() {
		var opt entity.OrderList
		err := rows.Scan(
			&opt.Total, &opt.ID, &opt.Ref, &opt.Customer.CustomerID, &opt.Customer.CustomerName, &opt.Customer.CustomerEmail, &opt.Customer.CustomerPhone, &opt.TotalQuantity, &opt.TotalPrice,
			&opt.CreatedAt, &opt.CreatedBy, &opt.UpdatedAt, &opt.UpdatedBy,
		)
		if err != nil {
			return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all orders")
		}
		orders = append(orders, &opt)
	}

	return orders, nil
}

func (r *repository) FindByID(ctx context.Context, ID string) (*entity.OrderList, error) {
	var opt entity.OrderList
	err := r.db.QueryRow(ctx, `
	SELECT 
		"order".id,
		"order".ref,
		customer.id as customer_id,
		customer.name as customer_name,
		customer.email as customer_email,
		customer.phone as customer_phone,
		"order".total_quantity,
		"order".total_price,
		"order".created_at,
		"order".created_by,
		"order".updated_at,
		"order".updated_by
	FROM "order" 
	JOIN customer ON "order".customer_id = customer.id
	WHERE "order".deleted_at IS NULL AND "order".id = $1`, ID).Scan(
		&opt.ID, &opt.Ref, &opt.Customer.CustomerID, &opt.Customer.CustomerName, &opt.Customer.CustomerEmail, &opt.Customer.CustomerPhone, &opt.TotalQuantity, &opt.TotalPrice,
		&opt.CreatedAt, &opt.CreatedBy, &opt.UpdatedAt, &opt.UpdatedBy,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, stacktrace.WrapWithCode(err, http.StatusNotFound, "order not found")
		}
		return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find order by id")
	}
	return &opt, nil
}

func (r *repository) Create(ctx context.Context, order entity.Order) (string, error) {
	query := `
	INSERT INTO "order" (ref, customer_id, total_quantity, total_price, created_by)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	err := r.db.QueryRow(ctx, query,
		order.Ref, order.CustomerID, order.TotalQuantity, order.TotalPrice,
		order.CreatedBy,
	).Scan(
		&order.ID,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return "", stacktrace.WrapWithCode(err, http.StatusConflict, "order code already exists")
		}

		return "", stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when create order")
	}
	return order.ID, nil
}

func (r *repository) Update(ctx context.Context, order entity.Order) (string, error) {
	query := `
        UPDATE "order"
        SET customer_id = $1, total_quantity = $2, total_price = $3, updated_at = $4, updated_by = $5
        WHERE id = $6
        RETURNING id`

	err := r.db.QueryRow(ctx, query,
		order.CustomerID, order.TotalQuantity, order.TotalPrice,
		order.UpdatedAt, order.UpdatedBy, order.ID,
	).Scan(
		&order.ID,
	)

	if err != nil {
		return "", stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when update order")
	}
	return order.ID, nil
}

func (r *repository) Delete(ctx context.Context, email, ID string) error {
	_, err := r.db.Exec(ctx, `UPDATE "order" set deleted_at = now(), deleted_by = $1 WHERE id = $2`,
		email, ID,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return stacktrace.WrapWithCode(err, http.StatusNotFound, "order not found")
		}
		return stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when delete order")
	}
	return nil
}
