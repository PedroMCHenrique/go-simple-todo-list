package config

import (
	"os"

	"github.com/pedromchenrique/todo-list/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeSqlite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	logger := GetLogger("main")

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database data not found")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite open error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Task{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
