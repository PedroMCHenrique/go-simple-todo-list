package main

import (
	"github.com/pedromchenrique/todo-list/config"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
}
