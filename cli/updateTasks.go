package cli

import (
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
)

type UpdateTasks struct{}

func (m model) updateTasks(msg tea.Msg) (tea.Model, tea.Cmd) {
	tasks := m.getTasks()
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
			if m.cursor < len(tasks)-1 {
				m.cursor++
			}

		case " ":
			if len(tasks) > 0 {
				tasks[m.cursor].IsDone = !tasks[m.cursor].IsDone
				m.tasksService.UpdateTask(tasks[m.cursor])
			}

		case "n", "N":
			m.viewContext = models.ViewNewTask
			m.textInput.SetValue("")
			m.textInput.Focus()

		case "e", "E":
			if len(tasks) > 0 {
				m.viewContext = models.ViewEditTask
				m.textInput.SetValue(tasks[m.cursor].Description)
				m.textInput.Focus()
			}
		case "d", "delete":
			if len(tasks) > 0 {
				m.tasksService.DeleteTask(tasks[m.cursor])
				m.cursor--

				if m.cursor < 0 {
					m.cursor = 0
				}

			}
		}
	case UpdateTasks:
	}
	return m, nil
}
