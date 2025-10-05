package main

import (
	"todo-cli/cli"
	"todo-cli/commands"
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
	command := commands.Init(db, context, config)
	if !command.IsCommand() {
		cli.Start(db, context, config)
	}
}
