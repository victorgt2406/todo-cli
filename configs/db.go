package configs

import (
	"fmt"
	"os"
	"todo-cli/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	SQLITE          = "sqlite"
	POSTGRES        = "postgres"
	MIGRATION_ERROR = "Error migrating DB: "
)

func InitDB() *gorm.DB {
	db_provider := CONFIG.Database.Provider
	switch db_provider {
	case SQLITE:
		return initSQLiteDB()
	case POSTGRES:
		return initPostgresDB()
	default:
		panic("Invalid database provider: " + db_provider)
	}
}

func initSQLiteDB() *gorm.DB {
	isNewDatabase := false
	if _, err := os.Stat(DB_PATH); os.IsNotExist(err) {
		isNewDatabase = true
	}
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})
	if err != nil {
		panic(MIGRATION_ERROR + err.Error())
	}
	if isNewDatabase {
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			panic(MIGRATION_ERROR + err.Error())
		}
		fmt.Printf("📂 Your sqlite database has been created and migrated! (%s)\n", DB_PATH)
	}
	return db
}

func initPostgresDB() *gorm.DB {
	isNewDatabase := false
	db, err := gorm.Open(postgres.Open(CONFIG.Database.Url), &gorm.Config{})
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
