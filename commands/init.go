package commands

import (
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
) Command {
	return Command{
		tasksService: tasksService.InitTaskService(db),
		llmService:   llmService.InitLlmService(config.LlmProvider),
		features:     config.Features,
		dbContext:    dbContext,
	}
}
