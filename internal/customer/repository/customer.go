package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/adopabianko/dbo-service/internal/customer/dto"
	"github.com/adopabianko/dbo-service/internal/customer/entity"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
)

func (r *repository) FindAll(ctx context.Context, params dto.CustomerListRequest) ([]entity.Customer, error) {
	query := `
        SELECT
 			COUNT(*) OVER () as total,
			id,
            name,
			phone,
			email,
			gender,
			address,
            created_at,
            created_by,
            updated_at,
            updated_by
        FROM customer
        WHERE deleted_at IS NULL`

	if params.Search != "" {
		query += fmt.Sprintf(` AND (name ILIKE '%s' OR phone ILIKE '%s' OR email ILIKE '%s')`, "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	query += fmt.Sprintf(` ORDER BY %s LIMIT %d OFFSET %d`, params.SortBy, params.Limit, (params.Page-1)*params.Limit)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all customers")
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var opt entity.Customer
		err := rows.Scan(
			&opt.Total, &opt.ID, &opt.Name, &opt.Phone, &opt.Email, &opt.Gender, &opt.Address,
			&opt.CreatedAt, &opt.CreatedBy, &opt.UpdatedAt, &opt.UpdatedBy,
		)
		if err != nil {
			return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all customers")
		}
		customers = append(customers, opt)
	}

	return customers, nil
}

func (r *repository) FindByID(ctx context.Context, ID string) (entity.Customer, error) {
	var opt entity.Customer
	err := r.db.QueryRow(ctx, `
	SELECT 
		id,
		name,
		phone,
		email,
		gender,
		address,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM customer WHERE deleted_at IS NULL AND id = $1`, ID).Scan(
		&opt.ID, &opt.Name, &opt.Phone, &opt.Email, &opt.Gender, &opt.Address,
		&opt.CreatedAt, &opt.CreatedBy, &opt.UpdatedAt, &opt.UpdatedBy,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusNotFound, "customer not found")
		}
		return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find customer by id")
	}
	return opt, nil
}

func (r *repository) Create(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	query := `
	INSERT INTO customer (name, phone, email, gender, address, created_by)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	err := r.db.QueryRow(ctx, query,
		customer.Name, customer.Phone, customer.Email, customer.Gender, customer.Address,
		customer.CreatedBy,
	).Scan(
		&customer.ID,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusConflict, "customer code already exists")
		}

		return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when create customer")
	}
	return customer, nil
}

func (r *repository) Update(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	query := `
        UPDATE customer
        SET name = $1, phone = $2, email = $3, gender = $4, address = $5, updated_at = now(), updated_by = $6
        WHERE id = $7
        RETURNING id`

	err := r.db.QueryRow(ctx, query,
		customer.Name, customer.Phone, customer.Email, customer.Gender, customer.Address,
		customer.UpdatedBy, customer.ID,
	).Scan(
		&customer.ID,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusConflict, "customer code already exists")
		}

		return entity.Customer{}, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when update customer")
	}
	return customer, nil
}

func (r *repository) Delete(ctx context.Context, email, ID string) error {
	_, err := r.db.Exec(ctx, `UPDATE customer set deleted_at = now(), deleted_by = $1 WHERE id = $2`,
		email, ID,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return stacktrace.WrapWithCode(err, http.StatusNotFound, "customer not found")
		}
		return stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when delete customer")
	}
	return nil
}
