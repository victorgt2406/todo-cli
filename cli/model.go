package cli

import (
	"todo-cli/db"
	"todo-cli/models"

	"github.com/charmbracelet/bubbles/textinput"
	"gorm.io/gorm"
)

type viewContext string

const (
	viewNewTask  viewContext = "newTask"
	viewEditTask viewContext = "editTask"
	viewTasks    viewContext = "tasks"
)

// The model is where the state of the cli is stored
type model struct {
	db          *gorm.DB
	dbContext   db.Context
	tasks       []models.Task
	cursor      int
	viewContext viewContext
	textInput   textinput.Model
}
