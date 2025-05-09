package repository

import (
	"context"
	"net/http"
	"strings"

	"github.com/adopabianko/dbo-service/internal/order/entity"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"
)

func (r *repository) FindOrderItemsByOrderId(ctx context.Context, orderID string) ([]entity.OrderItemList, error) {
	query := `
        SELECT
			order_item.id,
			order_item.order_id,
			product.id AS product_id,
			product.sku AS product_sku,
			product.name AS product_name,
			order_item.quantity,
			order_item.subtotal
		FROM order_item
		JOIN product ON order_item.product_id = product.id
		WHERE order_item.order_id = $1`

	rows, err := r.db.Query(ctx, query, orderID)
	if err != nil {
		return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all orderItems")
	}
	defer rows.Close()

	var orderItems []entity.OrderItemList
	for rows.Next() {
		var orderItem entity.OrderItemList
		err := rows.Scan(
			&orderItem.ID, &orderItem.OrderID, &orderItem.Product.ProductID, &orderItem.Product.ProductSku, &orderItem.Product.ProductName, &orderItem.Quantity, &orderItem.Subtotal,
		)
		if err != nil {
			return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find all orderItems")
		}
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (r *repository) CreateOrderItem(ctx context.Context, orderItem entity.OrderItem) (string, error) {
	query := `
	INSERT INTO order_item (order_id, product_id, quantity, subtotal)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	err := r.db.QueryRow(ctx, query,
		&orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Subtotal,
	).Scan(
		&orderItem.ID,
	)

	if err != nil {
		return "", stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when create order item")
	}
	return orderItem.ID, nil
}

func (r *repository) DeleteOrderItem(ctx context.Context, orderID string) error {
	query := `DELETE FROM "order_item" WHERE order_id in($1)`
	_, err := r.db.Exec(ctx, query, orderID)
	if err != nil {
		return stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when delete order item")
	}
	return nil
}

func (r *repository) FindProductByID(ctx context.Context, ID string) (*entity.Product, error) {
	var product entity.Product
	err := r.db.QueryRow(ctx, `
	SELECT
		id,
		sku,
		name,
		price
	FROM product WHERE id = $1`, ID).Scan(
		&product.ID, &product.Sku, &product.Name, &product.Price,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, stacktrace.WrapWithCode(err, http.StatusNotFound, "product not found")
		}
		return nil, stacktrace.WrapWithCode(err, http.StatusInternalServerError, "error when find product by id")
	}
	return &product, nil
}
