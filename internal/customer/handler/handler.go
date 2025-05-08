package handler

import (
	"github.com/adopabianko/dbo-service/internal/customer/service"

	"github.com/gin-gonic/gin"
)

type (
	ICustomerHandler interface {
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	handler struct {
		service service.ICustomerService
	}
)

func NewHandler(customerService service.ICustomerService) ICustomerHandler {
	return &handler{
		service: customerService,
	}
}
