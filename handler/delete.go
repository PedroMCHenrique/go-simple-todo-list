package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

func Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "id is required")
	}

	task := db.First(&schemas.Task{}, id).Error

	if task != nil {
		logger.Error("task not found")
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}

	err := db.Delete(&schemas.Task{}, id).Error
	if err != nil {
		logger.Error("task not found")
		SendError(ctx, http.StatusNotFound, "task not found")
		return
	}

	SendSuccess(ctx, "task deleted", task)
}
