package db

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgresDB() *gorm.DB {
	isNewDatabase := false
	db, err := gorm.Open(postgres.Open(configs.CONFIG.Database.Url), &gorm.Config{})
	if err != nil {
		panic(MIGRATION_ERROR + err.Error())
	}
	if !db.Migrator().HasTable(&models.Task{}) {
		isNewDatabase = true
	}
	if isNewDatabase {
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			panic(MIGRATION_ERROR + err.Error())
		}
		fmt.Println("🐘 Your PostgreSQL database has been created and migrated!")
	}

	return db
}
