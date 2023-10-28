package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

func CreateTask(ctx *gin.Context) {
	request := createTaskRequest{}
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		logger.Errorf("error binding request: %v", err)
		SendError(ctx, http.StatusBadRequest, "error binding request")
	}

	newTask := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
	}

	err = db.Create(&newTask).Error

	if err != nil {
		logger.Errorf("error creating task")
		SendError(ctx, http.StatusBadRequest, "error creating task")
	}

	SendSuccess(ctx, "create task", newTask)
}
