package cli

import (
	"todo-cli/config/configFile"
	"todo-cli/db"
	"todo-cli/models"
	tp "todo-cli/presenters/tasksPresenter"
	"todo-cli/services/llmService"
	"todo-cli/services/tasksService"

	"github.com/charmbracelet/bubbles/textinput"
)

// The model is where the state of the cli is stored
type model struct {
	// Services
	tasksService tasksService.TasksService
	llmService   *llmService.LlmService
	// Presenters
	tasksPresenter tp.TasksPresenter
	// Features
	features configFile.Features
	// UI
	cursor      int
	viewContext models.ViewContext
	textInput   textinput.Model
}

type initModelProps struct {
	tasksService tasksService.TasksService
	dbContext    db.Context
	llmService   *llmService.LlmService
	features     configFile.Features
}

func initModel(
	p initModelProps,
) model {
	textInput := textinput.New()
	textInput.Prompt = ""
	tasksPresenter := tp.InitTasksPresenter(p.dbContext)

	return model{
		tasksService:   p.tasksService,
		viewContext:    models.ViewTasks,
		textInput:      textInput,
		llmService:     p.llmService,
		features:       p.features,
		tasksPresenter: tasksPresenter,
	}
}
