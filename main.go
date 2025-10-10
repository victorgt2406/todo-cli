package main

import (
	"todo-cli/cli"
	"todo-cli/commands"
	"todo-cli/config/agentsMd"
	"todo-cli/config/configFile"
	"todo-cli/db"
	_ "todo-cli/internal/env"
)

func main() {
	db, dbContext := db.InitDb()
	config := configFile.LoadConfig()
	agentsMd := agentsMd.ReadAgentsMd()
	command := commands.Init(db, dbContext, config, agentsMd)
	if !command.IsCommand() {
		cli.Start(cli.TodoCliStartProps{
			Db:        db,
			DbContext: dbContext,
			Config:    config,
			AgentsMd:  agentsMd,
		})
	}
}
