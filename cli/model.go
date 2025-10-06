package cli

import (
	"todo-cli/config/configFile"
	"todo-cli/db"
	tasksPresenter "todo-cli/presenters/tasks"
	"todo-cli/services/llmService"
	"todo-cli/services/tasksService"

	"github.com/charmbracelet/bubbles/textinput"
)

type viewContext string

const (
	viewNewTask  viewContext = "newTask"
	viewEditTask viewContext = "editTask"
	viewTasks    viewContext = "tasks"
)

// The model is where the state of the cli is stored
type model struct {
	// Services
	tasksService tasksService.TasksService
	llmService   *llmService.LlmService
	// Presenters
	tasksPresenter tasksPresenter.TasksPresenter
	// Features
	features configFile.Features
	// UI
	cursor      int
	viewContext viewContext
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
	tasksPresenter := tasksPresenter.InitTasksPresenter(p.dbContext)

	return model{
		tasksService:   p.tasksService,
		viewContext:    viewTasks,
		textInput:      textInput,
		llmService:     p.llmService,
		features:       p.features,
		tasksPresenter: tasksPresenter,
	}
}
