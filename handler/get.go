package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary Get Task
// @Description Show Task detail
// @Tags tasks
// @Accept json
// @Produce json
// @Param id query string true "Task identification"
// @Success 200 {object} CreateTaskResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /task/:id [get]
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
