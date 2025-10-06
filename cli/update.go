package cli

import (
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	if m.viewContext == models.ViewNewTask || m.viewContext == models.ViewEditTask {
		return m.updateTask(msg)
	}
	return m.updateTasks(msg)
}
