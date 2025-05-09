package router

import (
	"github.com/adopabianko/dbo-service/internal/order/handler"
	"github.com/adopabianko/dbo-service/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

func Order(r *gin.Engine, orderHandler handler.IOrderHandler) {
	router := r.Group("/order")
	{
		router.Use(middleware.JWTMiddleware())
		router.GET("", orderHandler.FindAll)
		router.GET("/:id", orderHandler.FindByID)
		router.POST("", orderHandler.Create)
		router.PATCH("/:id", orderHandler.Update)
		router.DELETE("/:id", orderHandler.Delete)
	}
}
