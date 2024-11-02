package configs

import (
	"fmt"
	"os"
	"todo-cli/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_PATH = TODO_CLI_PATH + "todo-cli.db"

func InitDB() *gorm.DB {
	isNewDatabase := false
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		isNewDatabase = true
	}
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic("Error opening DB: " + err.Error())
	}
	if isNewDatabase {
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			panic("Error migrating DB: " + err.Error())
		}
		fmt.Printf("📂 DB created and migrated! (%s)\n", DB_PATH)
	}
	return db
}
