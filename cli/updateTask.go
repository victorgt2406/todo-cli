package cli

import (
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateTask(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.viewContext = viewTasks
			m.textInput.SetValue("")
			return m, nil
		case "enter":
			switch m.viewContext {
			case viewNewTask:
				return m.handleNewTask()
			case viewEditTask:
				return m.handleEditTask()
			}
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) handleNewTask() (tea.Model, tea.Cmd) {
	description := m.textInput.Value()
	if description != "" {
		// Create new task
		newTask := models.Task{
			Description: description,
			IsDone:      false,
		}

		// Save to database
		m.db.Create(&newTask)

		// Add to local tasks slice
		m.tasks = append(m.tasks, newTask)

		// Move cursor to the new task
		m.cursor = len(m.tasks) - 1
	}

	// Return to tasks view
	m.viewContext = viewTasks
	m.textInput.SetValue("")

	return m, nil
}

func (m model) handleEditTask() (tea.Model, tea.Cmd) {
	description := m.textInput.Value()
	if description != "" && len(m.tasks) > 0 {
		// Update task description
		m.tasks[m.cursor].Description = description

		// Save to database
		m.db.Save(&m.tasks[m.cursor])
	}

	// Return to tasks view
	m.viewContext = viewTasks
	m.textInput.SetValue("")

	return m, nil
}
