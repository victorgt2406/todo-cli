package main

import (
	"todo-cli/cli"
	"todo-cli/commands"
	"todo-cli/config/configFile"
	"todo-cli/db"
	_ "todo-cli/internal/env"
)

func main() {
	db, dbContext := db.InitDb()
	config := configFile.LoadConfig()
	command := commands.Init(db, dbContext, config)
	if !command.IsCommand() {
		cli.Start(cli.TodoCliStartProps{
			Db:        db,
			DbContext: dbContext,
			Config:    config,
		})
	}
}
