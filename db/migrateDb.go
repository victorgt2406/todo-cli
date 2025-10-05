package db

import (
	"fmt"
	"todo-cli/models"

	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB, dbPath string) {
	err := db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("Error when migrating: " + err.Error())
	}
	fmt.Printf("📂 Sqlite database created at\t(%s)\n", dbPath)
}
