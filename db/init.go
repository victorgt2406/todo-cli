package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	localDatabase := LocalDbPath()
	globalDatabase := GlobalDbPath()

	if dbExists(localDatabase) {
		return localDatabase, localStorage, false
	}
	if dbExists(globalDatabase) {
		return globalDatabase, globalStorage, false
	}
	return globalDatabase, globalStorage, true
}

func openDatabase(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Error when opening the database: " + err.Error())
	}
	return db
}
