package tasksPresenter

import (
	"fmt"
	"time"
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
	isCompletedTasks := false
	for i, task := range p.Tasks {
		description := task.Description
		if p.ViewContext == models.ViewEditTask && p.Cursor == i {
			description = p.TextInput.View()
		}
		if task.IsDone && !isCompletedTasks {
			s += "\nCompleted Tasks:\n"
			isCompletedTasks = true
		}
		s += t.viewTask(viewTaskProps{
			description: description,
			todoDate:    task.TodoDate,
			isDone:      task.IsDone,
			selected:    p.Cursor == i && p.ViewContext != models.ViewNewTask,
			viewContext: p.ViewContext,
		})
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

type viewTaskProps struct {
	description string
	todoDate    *time.Time
	isDone      bool
	selected    bool
	viewContext models.ViewContext
}

func (t TasksPresenter) viewTask(p viewTaskProps) string {
	s := ""

	strCursor := " "
	checked := " "
	description := p.description
	if p.selected {
		strCursor = ">"
	}

	if p.isDone {
		checked = "x"
	}

	s += fmt.Sprintf("%s [%s] %s", strCursor, checked, description)
	if p.todoDate != nil {
		s += fmt.Sprintf(" 📅 %s", utils.FormatDateToString(*p.todoDate))
	}

	return s
}
