package cli

import (
	"fmt"
	"os"
	"todo-cli/db"
	"todo-cli/services/tasksService"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
)

// Start the cli
func Start(db *gorm.DB, context db.Context) {
	p := tea.NewProgram(initialModel(db, context))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// Initial state of the cli
func initialModel(db *gorm.DB, context db.Context) model {
	textInput := textinput.New()
	textInput.Prompt = ""

	return model{
		db:          db,
		tasks:       tasksService.GetTasks(db),
		dbContext:   context,
		viewContext: viewTasks,
		textInput:   textInput,
	}
}

// not needed
func (m model) Init() tea.Cmd {
	return nil
}
