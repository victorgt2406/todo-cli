package cli

import (
	"todo-cli/models"
	s "todo-cli/services/tasksService"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateTask(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.viewContext = models.ViewTasks
			m.textInput.SetValue("")
			return m, nil
		case "enter":
			switch m.viewContext {
			case models.ViewNewTask:
				return m.handleNewTask()
			case models.ViewEditTask:
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
		isDone := false
		m.cursor = len(m.tasksService.GetTasks(s.TaskFilter{
			IsDone: &isDone,
		}, s.OrderBy{})) - 1

		if m.features.SmartTask {
			taskToAnalize = &newTask
		}
	}

	m.viewContext = models.ViewTasks
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
	tasks := m.getTasks()

	if description != "" && len(tasks) > 0 {
		tasks[m.cursor].Description = description
		updatedTask := m.tasksService.UpdateTask(tasks[m.cursor])

		if m.features.SmartTask {
			taskToAnalize = &updatedTask
		}
	}

	m.viewContext = models.ViewTasks
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
