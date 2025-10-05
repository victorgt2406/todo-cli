package main

import (
	"todo-cli/cli"
	"todo-cli/config/configFile"
	"todo-cli/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, context := db.InitDb()
	config := configFile.LoadConfig()
	cli.Start(db, context, config)
}
