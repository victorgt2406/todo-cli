package db

import (
	"fmt"
	"os"
	"path/filepath"

	"todo-cli/config"
	"todo-cli/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Context string

const (
	globalStorage Context = "global"
	localStorage  Context = "local"
)

const DATABASE_FILE_NAME = "todo-cli.db"

func InitDb() (*gorm.DB, Context) {
	dbPath, storageContext, isNewDatabase := whichDatabase()
	db := openDatabase(dbPath)

	if isNewDatabase {
		migrateDatabase(db, dbPath)
	}

	return db, storageContext
}

func whichDatabase() (string, Context, bool) {
	localDatabase := filepath.Join(config.GetLocalAppDir(), DATABASE_FILE_NAME)
	if _, err := os.Stat(localDatabase); err == nil {
		return localDatabase, localStorage, false
	}

	globalDatabase := filepath.Join(config.GetGlobalAppDir(), DATABASE_FILE_NAME)
	if _, err := os.Stat(globalDatabase); err == nil {
		return globalDatabase, globalStorage, false
	}

	return globalDatabase, globalStorage, true
}

func openDatabase(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Error when opening the database: " + err.Error())
	}
	return db
}

func migrateDatabase(db *gorm.DB, dbPath string) {
	err := db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("Error when migrating: " + err.Error())
	}
	fmt.Printf("📂 Sqlite database created at\t(%s)\n", dbPath)
}
