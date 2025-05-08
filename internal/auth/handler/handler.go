package handler

import (
	"github.com/adopabianko/dbo-service/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type (
	IAuthHandler interface {
		Login(ctx *gin.Context)
		Register(ctx *gin.Context)
	}

	handler struct {
		service service.IAuthService
	}
)

func NewHandler(authService service.IAuthService) IAuthHandler {
	return &handler{
		service: authService,
	}
}
