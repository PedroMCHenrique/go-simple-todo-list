package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3001")
}
