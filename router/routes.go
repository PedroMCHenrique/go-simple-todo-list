package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pedromchenrique/todo-list/docs"
	"github.com/pedromchenrique/todo-list/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	routes.Docs(v1)
}

func (r *routes) Task(rg *gin.RouterGroup) {
	task := rg.Group("/task")
	task.POST("/", handler.CreateTask)
	task.GET("/", handler.List)
	task.GET("/:id", handler.Get)
	task.PUT("/:id", handler.Update)
	task.DELETE("/:id", handler.Delete)
}

func (r *routes) Docs(rg *gin.RouterGroup) {
	docs := rg.Group("/docs")
	docs.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:3001/api/v1/docs/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}
