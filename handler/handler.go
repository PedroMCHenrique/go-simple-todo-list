package handler

import (
	"github.com/pedromchenrique/todo-list/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetDb()
}
