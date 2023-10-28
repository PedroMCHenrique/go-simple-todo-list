package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/schemas"
)

func CreateTask(ctx *gin.Context) {
	request := createTaskRequest{}
	err := ctx.ShouldBindJSON(&request)
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		logger.Errorf("error binding request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error binding request",
		})
		return
	}

	newTask := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
	}

	err = db.Create(&newTask).Error

	if err != nil {
		logger.Errorf("error creating task")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating task",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": newTask,
	})
}
