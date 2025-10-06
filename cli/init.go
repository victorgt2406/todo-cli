package cli

import (
	"fmt"
	"os"
	"todo-cli/config/configFile"
	"todo-cli/db"
	"todo-cli/services/llmService"
	"todo-cli/services/tasksService"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
)

// Start the cli
func Start(db *gorm.DB, context db.Context, config configFile.ConfigFile) {
	tasksService := tasksService.InitTaskService(db)
	llmService := llmService.InitLlmService(config.LlmProvider)

	p := tea.NewProgram(initialModel(tasksService, context, &llmService, config.Features))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// Initial state of the cli
func initialModel(
	tasksService tasksService.TasksService,
	dbContext db.Context,
	llmService *llmService.LlmService,
	features configFile.Features,
) model {
	textInput := textinput.New()
	textInput.Prompt = ""

	return model{
		tasksService: tasksService,
		dbContext:    dbContext,
		viewContext:  viewTasks,
		textInput:    textInput,
		llmService:   llmService,
		features:     features,
	}
}

// not needed
func (m model) Init() tea.Cmd {
	return nil
}
