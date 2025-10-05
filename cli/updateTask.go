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
	var taskToAnalize *models.Task = nil

	if description != "" {
		newTask := m.tasksService.CreateTask(description)
		m.tasks = append(m.tasks, newTask)
		m.cursor = len(m.tasks) - 1

		if m.features.SmartTask {
			taskToAnalize = &newTask
		}
	}

	m.viewContext = viewTasks
	m.textInput.SetValue("")

	return m, tea.Batch(
		func() tea.Msg {
			if taskToAnalize != nil {
				m.tasksService.UpdateTask(m.llmService.AnalizeTask(*taskToAnalize))
			}
			return UpdateTasks{}
		},
	)
}

func (m model) handleEditTask() (tea.Model, tea.Cmd) {
	description := m.textInput.Value()
	var taskToAnalize *models.Task = nil

	if description != "" && len(m.tasks) > 0 {
		m.tasks[m.cursor].Description = description
		updatedTask := m.tasksService.UpdateTask(m.tasks[m.cursor])

		if m.features.SmartTask {
			taskToAnalize = &updatedTask
		}
	}

	m.viewContext = viewTasks
	m.textInput.SetValue("")

	return m, tea.Batch(
		func() tea.Msg {
			if taskToAnalize != nil {
				m.tasksService.UpdateTask(m.llmService.AnalizeTask(*taskToAnalize))
			}
			return UpdateTasks{}
		},
	)
}
