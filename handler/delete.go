package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary Delete task
// @Description Delete a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id query string true "task identification"
// @Success 200 {object} CreateTaskResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /task [delete]
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
