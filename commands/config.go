package commands

import (
	"log"
	viewConfig "todo-cli/views/config"

	tea "github.com/charmbracelet/bubbletea"
)

func Config() {
	p := tea.NewProgram(viewConfig.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
