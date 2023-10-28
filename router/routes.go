package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromchenrique/todo-list/handler"
)

type routes struct {
	router *gin.Engine
}

func initializeRoutes(router *gin.Engine) {
	handler.Init()
	v1 := router.Group("/api/v1")

	routes := routes{
		router: router,
	}
	routes.Task(v1)
}

func (r *routes) Task(rg *gin.RouterGroup) {
	task := rg.Group("/task")
	task.POST("/", handler.CreateTask)
	task.GET("/", handler.List)
}
