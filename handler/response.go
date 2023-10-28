package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s", op),
		"data":    data,
	})
}

func SendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error": message,
	})
}
