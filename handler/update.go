package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

// @BasePath /api/v1

// @Summary Update task
// @Description Update a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id query string true "task Identification"
// @Param task body updateTaskRequest true "task data to Update"
// @Success 200 {object} CreateTaskResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /task [put]
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, "id is required")
	}

	request := updateTaskRequest{}
	err := ctx.BindJSON(&request)

	if err != nil {
		logger.Errorf("error binding request: %v", err)
		SendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	var task schemas.Task

	err = db.First(&task, id).Error

	if err != nil {
		logger.Errorf("error binding request: %v", err)
		SendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	task.Title = request.Title
	task.Description = request.Description
	task.Completed = request.Completed

	err = db.Save(&task).Error

	if err != nil {
		logger.Errorf("error binding request: %v", err)
		SendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	SendSuccess(ctx, "update task", task)
}
