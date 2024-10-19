package views

import (
	"fmt"
	"todo-cli/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func SelectTask(tasks []models.Task) (*models.Task, string, error) {
	p := tea.NewProgram(initialTaskModel(tasks))
	m, err := p.Run()
	if err != nil {
		return nil, "", fmt.Errorf("error running program: %v", err)
	}

	if m, ok := m.(taskModel); ok && m.selected != -1 {
		return &tasks[m.selected], m.command, nil
	}
	return nil, "", fmt.Errorf("no task selected")
}

type taskModel struct {
	tasks    []models.Task
	cursor   int
	selected int
	command  string
}

func initialTaskModel(tasks []models.Task) taskModel {
	return taskModel{
		tasks:    tasks,
		selected: -1,
	}
}

func (m taskModel) Init() tea.Cmd {
	return nil
}

func (m taskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.command = msg.String()
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
		case "enter", "e", "d": // list commands
			m.selected = m.cursor
			fmt.Println("")
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m taskModel) View() string {
	s := "Select a task:\n\n"

	for i, task := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		status := "[-]"
		description := task.Description
		if task.IsDone {
			status = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("[x]")
			description = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(description)
		}
		s += fmt.Sprintf("%s %s %s\n", cursor, status, description)
	}

	s += "\nPress q to quit.\n"
	return s
}
