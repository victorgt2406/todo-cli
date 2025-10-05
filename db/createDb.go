package db

import (
	"path/filepath"
	"todo-cli/config"
)

func CreateDb(dbPath string) {
	if dbExists(dbPath) {
		panic("Database already exists: " + dbPath)
	}

	db, err := openDb(dbPath)
	if err != nil {
		panic("Error when opening the database: " + err.Error())
	}

	migrateDatabase(db, dbPath)
}

func LocalDbPath() string {
	return filepath.Join(config.GetLocalAppDir(), DATABASE_FILE_NAME)
}

func GlobalDbPath() string {
	return filepath.Join(config.GetGlobalAppDir(), DATABASE_FILE_NAME)
}
