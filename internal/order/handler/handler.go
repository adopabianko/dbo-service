package handler

import (
	"github.com/adopabianko/dbo-service/internal/order/service"
	"github.com/gin-gonic/gin"
)

type (
	IOrderHandler interface {
		FindAll(ctx *gin.Context)
		FindByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	handler struct {
		service service.IOrderService
	}
)

func NewHandler(orderService service.IOrderService) IOrderHandler {
	return &handler{
		service: orderService,
	}
}
