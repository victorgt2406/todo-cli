package db

import (
	"todo-cli/configs"

	"gorm.io/gorm"
)

const (
	SQLITE          = "sqlite"
	POSTGRES        = "postgres"
	MIGRATION_ERROR = "Error migrating DB: "
)

func InitDB() *gorm.DB {
	db_provider := configs.CONFIG.Database.Provider
	switch db_provider {
	case SQLITE:
		return initSQLiteDB()
	case POSTGRES:
		return initPostgresDB()
	default:
		panic("Invalid database provider: " + db_provider)
	}
}
