package cli

import (
	"fmt"
	"os"
	"todo-cli/config/db"
	"todo-cli/services/llm"
	"todo-cli/services/tasksService"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"gorm.io/gorm"
)

// Start the cli
func Start(db *gorm.DB, context db.Context) {
	tasksService := tasksService.InitTaskService(db)
	llmService := llm.InitLlmService()

	p := tea.NewProgram(initialModel(tasksService, context, llmService))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// Initial state of the cli
func initialModel(
	tasksService tasksService.TasksService,
	dbContext db.Context,
	llmService llm.LlmService,
) model {
	textInput := textinput.New()
	textInput.Prompt = ""

	return model{
		tasksService: tasksService,
		tasks:        tasksService.GetTasks(),
		dbContext:    dbContext,
		viewContext:  viewTasks,
		textInput:    textInput,
		llmService:   llmService,
	}
}

// not needed
func (m model) Init() tea.Cmd {
	return nil
}
