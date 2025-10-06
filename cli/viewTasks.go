package cli

import (
	"fmt"
	"todo-cli/utils"
)

func (m model) viewTasks() string {
	// The header
	s := "Todo List\n\n"

	tasks := m.getTasks()

	// Tasks list
	for i, task := range tasks {
		cursor := " "  // no cursor
		checked := " " // not selected
		description := task.Description
		todoDate := utils.FormatDateToString(task.TodoDate)
		selected := m.cursor == i && m.viewContext != viewNewTask

		if selected {
			cursor = ">" // cursor!
			if m.viewContext == viewEditTask {
				description = m.textInput.View()
			}
		}

		if task.IsDone {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s", cursor, checked, description)
		if todoDate != nil {
			s += fmt.Sprintf(" 📅 %s", *todoDate)
		}
		s += "\n"
	}

	if m.viewContext == viewNewTask {
		m.textInput.Prompt = "> [ ] "
		s += m.textInput.View() + "\n"

	}

	// The footer
	s += "\nPress q to quit. Press n to add a new task. Press e to edit selected task. Press space to toggle completion.\n"

	// Send the UI for rendering
	return s
}
