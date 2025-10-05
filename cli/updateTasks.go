package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

type UpdateTasks struct{}

func (m model) updateTasks(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "Q":
			return m, tea.Quit

		case "up", "k", "K":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j", "J":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		case " ":
			if len(m.tasks) > 0 {
				m.tasks[m.cursor].IsDone = !m.tasks[m.cursor].IsDone
				m.tasksService.UpdateTask(m.tasks[m.cursor])
			}

		case "n", "N":
			m.viewContext = viewNewTask
			m.textInput.SetValue("")
			m.textInput.Focus()

		case "e", "E":
			if len(m.tasks) > 0 {
				m.viewContext = viewEditTask
				m.textInput.SetValue(m.tasks[m.cursor].Description)
				m.textInput.Focus()
			}
		case "d", "delete":
			if len(m.tasks) > 0 {
				m.tasksService.DeleteTask(m.tasks[m.cursor])
				m.cursor--
				m.tasks = m.tasksService.GetTasks()

				if m.cursor < 0 {
					m.cursor = 0
				}

			}
		}
	case UpdateTasks:
		m.tasks = m.tasksService.GetTasks()
	}
	return m, nil
}
