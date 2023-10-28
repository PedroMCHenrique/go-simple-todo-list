package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

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
