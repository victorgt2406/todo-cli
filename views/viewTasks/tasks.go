package viewTasks

import (
	"fmt"
	"slices"
	"strings"
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Tasks(tasks []models.Task, allowedCommands []string) (*models.Task, string, error) {
	p := tea.NewProgram(initialTeaModel(tasks, allowedCommands))
	m, err := p.Run()

	if err != nil {
		return nil, "", fmt.Errorf("error running program: %v", err)
	}

	taskModel := m.(teaModel)
	if taskModel.selected == -1 {
		return nil, "", fmt.Errorf("no task selected")
	}

	return &taskModel.tasks[taskModel.selected], taskModel.command, nil

}

type teaModel struct {
	tasks           []models.Task
	cursor          int
	selected        int
	allowedCommands []string
	command         string
}

func initialTeaModel(tasks []models.Task, allowedCommands []string) teaModel {
	return teaModel{
		tasks:           tasks,
		allowedCommands: allowedCommands,
		selected:        -1,
	}
}

func (m teaModel) Init() tea.Cmd {
	return nil
}

func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.command = msg.String()
		if slices.Contains(m.allowedCommands, m.command) {
			m.selected = m.cursor
			fmt.Println("Command: "+m.command, "Commands: "+strings.Join(m.allowedCommands, ", "))
			return m, tea.Quit
		} else {
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}
			}
		}
	}
	return m, nil
}

func (m teaModel) View() string {
	s := "Select a task:\n\n"

	for i, task := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		status := "[-]"
		description := task.Description
		date := task.Date
		if task.IsDone {
			status = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("[x]")
			description = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(description)
		}
		s += fmt.Sprintf("%s %s %s %s\n", cursor, status, description, date)
	}

	s += "\nPress q to quit.\n"
	return s
}
