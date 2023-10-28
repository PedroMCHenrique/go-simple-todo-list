package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
