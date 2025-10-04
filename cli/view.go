package cli

import "fmt"

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
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, description)
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
