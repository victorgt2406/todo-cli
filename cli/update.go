package cli

import (
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle text input updates first if we're in input mode
	if m.viewContext == viewNewTask || m.viewContext == viewEditTask {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)

		// Check for special keys in input mode
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			switch keyMsg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
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

		return m, cmd
	}

	// Handle navigation in tasks view
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// Navigation keys
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		// Toggle task completion with spacebar
		case " ":
			if len(m.tasks) > 0 {
				// Toggle the IsDone status
				m.tasks[m.cursor].IsDone = !m.tasks[m.cursor].IsDone
				// Save to database
				m.db.Save(&m.tasks[m.cursor])
			}

		case "n":
			m.viewContext = viewNewTask
			m.textInput.SetValue("")
			m.textInput.Focus()

		case "e":
			if len(m.tasks) > 0 {
				m.viewContext = viewEditTask
				m.textInput.SetValue(m.tasks[m.cursor].Description)
				m.textInput.Focus()
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
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
