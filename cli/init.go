package cli

import (
	"fmt"
	"os"
	"todo-cli/config/configFile"
	"todo-cli/db"
	"todo-cli/services/llmService"
	"todo-cli/services/tasksService"

	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
)

type TodoCliStartProps struct {
	Db        *gorm.DB
	DbContext db.Context
	Config    configFile.ConfigFile
}

// Start the cli
func Start(p TodoCliStartProps) {
	tasksService := tasksService.InitTaskService(p.Db)
	llmService := llmService.InitLlmService(p.Config.LlmProvider)

	cli := tea.NewProgram(initModel(initModelProps{
		tasksService: tasksService,
		dbContext:    p.DbContext,
		llmService:   &llmService,
		features:     p.Config.Features,
	}))

	if _, err := cli.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// not needed
func (m model) Init() tea.Cmd {
	return nil
}
