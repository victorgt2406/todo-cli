package main

import (
	"todo-cli/cli"
	"todo-cli/commands"
	"todo-cli/config/configFile"
	"todo-cli/db"
	_ "todo-cli/internal/env"
)

func main() {
	db, context := db.InitDb()
	config := configFile.LoadConfig()
	command := commands.Init(db, context, config)
	if !command.IsCommand() {
		cli.Start(db, context, config)
	}
}
