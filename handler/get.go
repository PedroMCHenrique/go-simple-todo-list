package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

func Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "id is required")
	}

	task := schemas.Task{}

	err := db.Where("id=?", id).First(&task).Error

	if err != nil {
		logger.Errorf("task not found: %v", err)
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}

	SendSuccess(ctx, "get task", task)
}
