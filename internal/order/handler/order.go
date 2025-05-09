package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adopabianko/dbo-service/internal/order/dto"
	"github.com/adopabianko/dbo-service/pkg/http/response"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"

	"github.com/gin-gonic/gin"
)

// @Tags Order
// @Summary Find All Order
// @Description API for find all order
// @Accept json
// @Produce json
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Param sort_by query string false "Sort By" default(created_at DESC)
// @Param search query string false "Search"
// @Param type query string false "Type"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order [get]
func (h *handler) FindAll(ctx *gin.Context) {
	var params dto.OrderListRequest

	// Get query parameters
	params.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	params.Limit, _ = strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	params.SortBy = ctx.DefaultQuery("sort_by", "created_at DESC")
	params.Search = ctx.DefaultQuery("search", "")

	// Call service layer
	result, err := h.service.FindAll(ctx, params)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "FindAll"),
			"failed to fetch orders",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"fetched all orders successfully",
		result.Data,
		result.Meta,
	)
}

// @Tags Order
// @Summary Get Order Detail By UUID
// @Description API for get order detail by uuid
// @Produce json
// @Param uuid path string true "order uuid"
// @Success 200 {object} response.SuccessResponse{data=entity.Order}
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order/{uuid} [get]
func (h *handler) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := h.service.FindByID(ctx, id)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "FindByID"),
			"failed to fetch order",
		)
		return
	}

	response.ResponseSuccess(
		ctx, http.StatusOK,
		"fetched order successfully",
		order,
		nil,
	)
}

// @Tags Order
// @Summary Create Order
// @Description API for create order
// @Accept json
// @Produce json
// @Param payload body dto.CreateOrderRequest true "Payload Create Order"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order [post]
func (h *handler) Create(ctx *gin.Context) {
	var order dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, http.StatusBadRequest, "Create"),
			"failed to create order",
		)
		return
	}

	// get email from token
	email, _ := ctx.Get("email")

	order.CreatedBy = email.(string)

	orderID, err := h.service.Create(ctx, order)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Create"),
			"failed to create order",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusCreated,
		"order created successfully",
		orderID,
		nil,
	)
}

// @Tags Order
// @Summary Create Order
// @Description API for create order
// @Accept json
// @Produce json
// @Param uuid path string true "order uuid"
// @Param payload body dto.UpdateOrderRequest true "Payload Update Order"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order/{uuid} [patch]
func (h *handler) Update(ctx *gin.Context) {
	var order dto.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&order); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Update"),
			"failed to update order",
		)
		return
	}

	// get email from token
	email, _ := ctx.Get("email")

	order.ID = ctx.Param("id")
	order.UpdatedAt = time.Now()
	order.UpdatedBy = email.(string)

	orderID, err := h.service.Update(ctx, order)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Update"),
			"failed to update order",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"order updated successfully",
		orderID,
		nil,
	)
}

// @Tags Order
// @Summary Delete Order
// @Description API for delete order
// @Produce json
// @Param uuid path string true "order uuid"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order/{uuid} [delete]
func (h *handler) Delete(ctx *gin.Context) { // get email from token
	email, _ := ctx.Get("email")
	id := ctx.Param("id")
	err := h.service.Delete(ctx, email.(string), id)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Delete"),
			"failed to delete order",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"order deleted successfully",
		nil,
		nil,
	)
}
