package commands

import (
	"todo-cli/config/agentsMd"
	"todo-cli/config/configFile"
	"todo-cli/db"
	"todo-cli/services/llmService"
	"todo-cli/services/tasksService"

	"gorm.io/gorm"
)

type Command struct {
	tasksService tasksService.TasksService
	llmService   llmService.LlmService
	features     configFile.Features
	dbContext    db.Context
}

func Init(
	db *gorm.DB,
	dbContext db.Context,
	config configFile.ConfigFile,
	agentsMd agentsMd.AgentsMd,
) Command {
	return Command{
		tasksService: tasksService.InitTaskService(db),
		llmService: llmService.InitLlmService(llmService.InitLlmServiceProps{
			LlmProvider: config.LlmProvider,
			AgentsMd:    agentsMd,
		}),
		features:  config.Features,
		dbContext: dbContext,
	}
}
