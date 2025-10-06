package tasksPresenter

import (
	"fmt"
	"todo-cli/models"
	"todo-cli/utils"

	"github.com/charmbracelet/bubbles/textinput"
)

type ViewTasksProps struct {
	Tasks       []models.Task
	Cursor      int
	ViewContext models.ViewContext
	TextInput   textinput.Model
}

func (t TasksPresenter) ViewTasks(p ViewTasksProps) string {
	// The header
	s := "Todo List\n\n"

	// Tasks list
	for i, task := range p.Tasks {
		strCursor := " " // no cursor
		checked := " "   // not selected
		description := task.Description
		todoDate := utils.FormatDateToString(task.TodoDate)
		selected := p.Cursor == i && p.ViewContext != models.ViewNewTask

		if selected {
			strCursor = ">" // cursor!
			if p.ViewContext == models.ViewEditTask {
				description = p.TextInput.View()
			}
		}

		if task.IsDone {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s", strCursor, checked, description)
		if todoDate != nil {
			s += fmt.Sprintf(" 📅 %s", *todoDate)
		}
		s += "\n"
	}

	if p.ViewContext == models.ViewNewTask {
		p.TextInput.Prompt = "> [ ] "
		s += p.TextInput.View() + "\n"

	}

	// The footer
	s += "\nPress q to quit. Press n to add a new task. Press e to edit selected task. Press space to toggle completion.\n"

	// Send the UI for rendering
	return s
}
