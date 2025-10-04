package cli

import (
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
	if m.viewContext == viewNewTask || m.viewContext == viewEditTask {
		return m.updateTaskView(msg)
	}
	return m.updateTasksView(msg)
}
