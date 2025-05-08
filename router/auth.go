package router

import (
	"github.com/adopabianko/dbo-service/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine, authHandler handler.IAuthHandler) {
	router := r.Group("/auth")
	{
		router.POST("login", authHandler.Login)
		router.POST("register", authHandler.Register)
	}
}
