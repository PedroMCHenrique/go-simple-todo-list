package main

import (
	"github.com/pedromchenrique/todo-list/config"
	"github.com/pedromchenrique/todo-list/router"
)

// @title           Simple Todo-list API
// @version         1.0
// @description     This is a simple Todo-list api.

// @host      localhost:3001
// @BasePath  /api/v1
func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initializing error: %v", err)
		return
	}

	router.Init()
}
