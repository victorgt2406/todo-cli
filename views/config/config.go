package config

import (
	"fmt"
	"todo-cli/configs"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
}

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m model) View() string {
	s := fmt.Sprintf("You can edit the config file at: %s\n", configs.CONFIG_PATH)
	return s
}
