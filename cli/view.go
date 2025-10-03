package cli

import "fmt"

func (m model) View() string {
	switch m.viewContext {
	case viewNewTask:
		return m.viewNewTask()
	case viewEditTask:
		return m.viewEditTask()
	default:
		return m.viewTasks()
	}
}

func (m model) viewTasks() string {
	// The header
	s := "Todo List\n\n"

	// Iterate over our choices
	for i, task := range m.tasks {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if task.IsDone {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Description)
	}

	// The footer
	s += "\nPress q to quit. Press n to add a new task. Press e to edit selected task. Press space to toggle completion.\n"

	// Send the UI for rendering
	return s
}

func (m model) viewNewTask() string {
	s := "Add New Task\n\n"
	s += "Enter task description:\n"
	s += m.textInput.View()
	s += "\n\nPress Enter to save, Esc to cancel.\n"
	return s
}

func (m model) viewEditTask() string {
	s := "Edit Task\n\n"
	s += "Enter new description:\n"
	s += m.textInput.View()
	s += "\n\nPress Enter to save, Esc to cancel.\n"
	return s
}
