package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/adopabianko/dbo-service/internal/order/dto"
	"github.com/adopabianko/dbo-service/internal/order/entity"
	"github.com/adopabianko/dbo-service/pkg/http/response"
)

func (s *service) FindAll(ctx context.Context, params dto.OrderListRequest) (dto.OrderListResponse, error) {
	var result dto.OrderListResponse

	orders, err := s.repository.FindAll(ctx, params)
	if err != nil {
		return result, err
	}

	var orderList []dto.Order
	for _, val := range orders {
		var order dto.Order

		customer, err := s.customerRepository.FindByID(ctx, val.Customer.CustomerID)
		if err != nil {
			return result, err
		}

		orderItems, err := s.repository.FindOrderItemsByOrderId(ctx, val.ID)
		if err != nil {
			return result, err
		}

		var quantity, totalPrice float32
		var orderItemsList []entity.OrderItemList
		for _, item := range orderItems {
			orderItemsList = append(orderItemsList, entity.OrderItemList{
				ID:       item.ID,
				Product:  item.Product,
				Quantity: item.Quantity,
				Subtotal: item.Subtotal,
			})

			quantity += item.Quantity
			totalPrice += item.Subtotal
		}

		order.ID = val.ID
		order.Ref = val.Ref
		order.Customer = entity.OrderCustomer{
			CustomerID:    customer.ID,
			CustomerName:  customer.Name,
			CustomerEmail: customer.Email,
			CustomerPhone: customer.Phone,
		}
		order.Items = orderItemsList
		order.CreatedAt = val.CreatedAt
		order.CreatedBy = val.CreatedBy
		order.UpdatedAt = val.UpdatedAt
		order.UpdatedBy = val.UpdatedBy
		order.TotalQuantity = quantity
		order.TotalPrice = totalPrice

		orderList = append(orderList, order)
	}

	var totalItems int
	if len(orders) > 0 {
		totalItems = orders[0].Total
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	result.Data = orderList
	result.Meta = response.BuildMeta(
		params.Limit,
		totalItems,
		params.Page,
		totalPages,
		params.SortBy,
	)

	return result, nil
}

func (s *service) FindByID(ctx context.Context, id string) (dto.Order, error) {
	var order dto.Order

	orderList, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return order, err
	}
	customer, err := s.customerRepository.FindByID(ctx, orderList.Customer.CustomerID)
	if err != nil {
		return order, err
	}
	orderItems, err := s.repository.FindOrderItemsByOrderId(ctx, orderList.ID)
	if err != nil {
		return order, err
	}
	var orderItemsList []entity.OrderItemList
	for _, item := range orderItems {
		orderItemsList = append(orderItemsList, entity.OrderItemList{
			ID:       item.ID,
			Product:  item.Product,
			Quantity: item.Quantity,
			Subtotal: item.Subtotal,
		})
	}
	order.ID = orderList.ID
	order.Ref = orderList.Ref
	order.Customer = entity.OrderCustomer{
		CustomerID:    customer.ID,
		CustomerName:  customer.Name,
		CustomerEmail: customer.Email,
		CustomerPhone: customer.Phone,
	}
	order.Items = orderItemsList
	order.CreatedAt = orderList.CreatedAt
	order.CreatedBy = orderList.CreatedBy
	order.UpdatedAt = orderList.UpdatedAt
	order.UpdatedBy = orderList.UpdatedBy
	order.TotalQuantity = orderList.TotalQuantity
	order.TotalPrice = orderList.TotalPrice
	return order, nil
}

func (s *service) Create(ctx context.Context, params dto.CreateOrderRequest) (dto.Response, error) {
	var order entity.Order

	var totalQuantity, totalPrice float32
	for _, item := range params.OrderItems {
		product, err := s.repository.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return dto.Response{}, err
		}

		totalPrice += product.Price * item.Quantity
		totalQuantity += item.Quantity
	}

	order.Ref = s.generateOrderRef()
	order.TotalQuantity = totalQuantity
	order.TotalPrice = totalPrice
	order.CustomerID = params.CustomerID
	order.CreatedBy = params.CreatedBy

	orderID, err := s.repository.Create(ctx, order)
	if err != nil {
		return dto.Response{}, err
	}

	for _, item := range params.OrderItems {
		product, err := s.repository.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return dto.Response{}, err
		}

		orderItem := entity.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Subtotal:  item.Quantity * product.Price,
		}
		_, err = s.repository.CreateOrderItem(ctx, orderItem)
		if err != nil {
			return dto.Response{}, err
		}
	}

	return dto.Response{ID: orderID}, nil
}

func (s *service) Update(ctx context.Context, params dto.UpdateOrderRequest) (dto.Response, error) {
	var order entity.Order

	orderList, err := s.repository.FindByID(ctx, params.ID)
	if err != nil {
		return dto.Response{}, err
	}

	var totalQuantity, totalPrice float32
	for _, item := range params.OrderItems {
		product, err := s.repository.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return dto.Response{}, err
		}

		totalPrice += product.Price * item.Quantity
		totalQuantity += item.Quantity
	}

	order.ID = orderList.ID
	order.UpdatedBy = &params.UpdatedBy
	order.UpdatedAt = &params.UpdatedAt
	order.TotalQuantity = totalQuantity
	order.TotalPrice = totalPrice
	order.CustomerID = params.CustomerID

	orderID, err := s.repository.Update(ctx, order)
	if err != nil {
		return dto.Response{}, err
	}

	// Delete order items that are not in the updated order items
	for _, _ = range params.OrderItems {
		err = s.repository.DeleteOrderItem(ctx, orderID)
		if err != nil {
			return dto.Response{}, err
		}
	}

	// Create or update order items
	for _, item := range params.OrderItems {
		product, err := s.repository.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return dto.Response{}, err
		}

		orderItem := entity.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Subtotal:  item.Quantity * product.Price,
		}

		_, err = s.repository.CreateOrderItem(ctx, orderItem)
		if err != nil {
			return dto.Response{}, err
		}
	}

	return dto.Response{ID: orderID}, nil
}

func (s *service) Delete(ctx context.Context, email, id string) error {
	return s.repository.Delete(ctx, email, id)
}

func (s *service) generateOrderRef() string {
	// Timestamp: format YYYYMMDDTHHMMSS
	timestamp := time.Now().Format("20060102T150405")

	// Random alphanumeric string (uppercase) with length 6
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(6)
	for range 6 {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	randomStr := sb.String()

	// Combine parts
	return fmt.Sprintf("ORD-%s-%s", timestamp, randomStr)
}
