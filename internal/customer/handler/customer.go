package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adopabianko/dbo-service/internal/customer/dto"
	"github.com/adopabianko/dbo-service/pkg/http/response"
	"github.com/adopabianko/dbo-service/pkg/stacktrace"

	"github.com/gin-gonic/gin"
)

// @Tags Customer
// @Summary Find All Customer
// @Description API for find all customer
// @Accept json
// @Produce json
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Param sort_by query string false "Sort By" default(created_at DESC)
// @Param search query string false "Search"
// @Param type query string false "Type"
// @Success 200 {object} response.SuccessResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /customer [get]
func (h *handler) FindAll(ctx *gin.Context) {
	var params dto.CustomerListRequest

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
			"failed to fetch customers",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"fetched all customers successfully",
		result.Data,
		result.Meta,
	)
}

// @Tags Customer
// @Summary Get Customer Detail By UUID
// @Description API for get customer detail by uuid
// @Produce json
// @Param uuid path string true "customer uuid"
// @Success 200 {object} response.SuccessResponse{data=entity.Customer}
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /customer/{uuid} [get]
func (h *handler) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := h.service.FindByID(ctx, id)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "FindByID"),
			"failed to fetch customer",
		)
		return
	}

	response.ResponseSuccess(
		ctx, http.StatusOK,
		"fetched customer successfully",
		customer,
		nil,
	)
}

// @Tags Customer
// @Summary Create Customer
// @Description API for create customer
// @Accept json
// @Produce json
// @Param payload body dto.CreateCustomerRequest true "Payload Create Customer"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /customer [post]
func (h *handler) Create(ctx *gin.Context) {
	var opt dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&opt); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, http.StatusBadRequest, "Create"),
			"failed to create customer",
		)
		return
	}

	// get email from token
	email, _ := ctx.Get("email")

	opt.CreatedBy = email.(string)

	createdCustomer, err := h.service.Create(ctx, opt)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Create"),
			"failed to create customer",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusCreated,
		"customer created successfully",
		createdCustomer,
		nil,
	)
}

// @Tags Customer
// @Summary Create Customer
// @Description API for create customer
// @Accept json
// @Produce json
// @Param uuid path string true "customer uuid"
// @Param payload body dto.UpdateCustomerRequest true "Payload Update Customer"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /customer/{uuid} [patch]
func (h *handler) Update(ctx *gin.Context) {
	var customer dto.UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Update"),
			"failed to update customer",
		)
		return
	}

	// get email from token
	email, _ := ctx.Get("email")

	customer.ID = ctx.Param("id")
	customer.UpdatedAt = time.Now()
	customer.UpdatedBy = email.(string)

	updatedCustomer, err := h.service.Update(ctx, customer)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Update"),
			"failed to update customer",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"customer updated successfully",
		updatedCustomer,
		nil,
	)
}

// @Tags Customer
// @Summary Delete Customer
// @Description API for delete customer
// @Produce json
// @Param uuid path string true "customer uuid"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /customer/{uuid} [delete]
func (h *handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.service.Delete(ctx, id)
	if err != nil {
		response.ResponseFailed(
			ctx,
			stacktrace.WrapWithCode(err, stacktrace.GetCode(err), "Delete"),
			"failed to delete customer",
		)
		return
	}

	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"customer deleted successfully",
		nil,
		nil,
	)
}
