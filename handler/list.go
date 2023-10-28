package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary List Tasks
// @Description List all tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} ListTaskResponse
// @Failure 500 {object} ErrorResponse
// @Router /task [get]
func List(ctx *gin.Context) {
	tasks := []schemas.Task{}
	err := db.Find(&tasks).Error

	if err != nil {
		logger.Errorf("error getting tasks: %v", err)
		SendError(ctx, http.StatusInternalServerError, "error getting tasks")
	}

	SendSuccess(ctx, "listing tasks", tasks)
}
