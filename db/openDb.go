package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDb(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	return db, err
}
