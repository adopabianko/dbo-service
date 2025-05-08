package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(r *gin.Engine) {
	router := r.Group("/")
	{
		router.GET("", Ping)
	}
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
