package cli

import (
	"fmt"
	"time"
)

func (m model) View() string {
	switch m.viewContext {
	default:
		return m.viewTasks()
	}
}

func (m model) viewTasks() string {
	// The header
	s := "Todo List\n\n"

	// Iterate over our choices
	for i, task := range m.tasks {
		cursor := " "  // no cursor
		checked := " " // not selected
		description := task.Description
		todoDate := formatDateToString(task.TodoDate)
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

func formatDateToString(t *time.Time) *string {
	if t == nil {
		return nil
	}
	weekday := t.Weekday().String()
	day := t.Day()
	month := t.Month().String()
	result := fmt.Sprintf("%s %d of %s", weekday, day, month)
	return &result
}
