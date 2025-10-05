package cli

import (
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
		newTask := m.tasksService.CreateTask(description)
		newTask = m.tasksService.UpdateTask(m.llmService.AnalizeTask(newTask))
		m.tasks = append(m.tasks, newTask)
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
		m.tasks[m.cursor].Description = description
		m.tasks[m.cursor] = m.llmService.AnalizeTask(m.tasks[m.cursor])
		m.tasks[m.cursor] = m.tasksService.UpdateTask(m.tasks[m.cursor])
	}

	// Return to tasks view
	m.viewContext = viewTasks
	m.textInput.SetValue("")

	return m, nil
}
