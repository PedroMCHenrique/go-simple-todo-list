package main

import (
	"github.com/pedromchenrique/todo-list/config"
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initializing error: %v", err)
		return
	}
}
