package router

import (
	"github.com/adopabianko/dbo-service/internal/customer/handler"
	"github.com/adopabianko/dbo-service/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

func Customer(r *gin.Engine, customerHandler handler.ICustomerHandler) {
	router := r.Group("/customer")
	{
		router.Use(middleware.JWTMiddleware())
		router.GET("", customerHandler.FindAll)
		router.GET("/:id", customerHandler.FindByID)
		router.POST("", customerHandler.Create)
		router.PATCH("/:id", customerHandler.Update)
		router.DELETE("/:id", customerHandler.Delete)
	}
}
