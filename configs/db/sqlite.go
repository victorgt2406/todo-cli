package db

import (
	"fmt"
	"os"
	"todo-cli/configs"
	"todo-cli/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSQLiteDB() *gorm.DB {
	isNewDatabase := false
	if _, err := os.Stat(configs.DB_PATH); os.IsNotExist(err) {
		isNewDatabase = true
	}
	db, err := gorm.Open(sqlite.Open(configs.DB_PATH), &gorm.Config{})
	if err != nil {
		panic(MIGRATION_ERROR + err.Error())
	}
	if isNewDatabase {
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			panic(MIGRATION_ERROR + err.Error())
		}
		fmt.Printf("📂 Your sqlite database has been created and migrated! (%s)\n", configs.DB_PATH)
	}
	return db
}
