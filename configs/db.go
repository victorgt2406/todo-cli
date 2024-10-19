package configs

import (
	"fmt"
	"os"
	"todo-cli/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbPath := "./todo-cli.db"
	isNewDatabase := false
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		isNewDatabase = true
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Error opening DB: " + err.Error())
	}
	if isNewDatabase {
		fmt.Println("Setting up DB")
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			panic("Error migrating DB: " + err.Error())
		}
		fmt.Println("📂 DB created and migrated!")
	}
	return db
}
