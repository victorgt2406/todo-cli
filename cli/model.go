package cli

import (
	"todo-cli/db"
	"todo-cli/models"
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
	tasksService tasksService.TasksService
	dbContext    db.Context
	tasks        []models.Task
	cursor       int
	viewContext  viewContext
	textInput    textinput.Model
}
