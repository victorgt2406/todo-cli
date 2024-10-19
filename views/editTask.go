package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type editTask struct {
	textInput textinput.Model
	task      string
}

func InitEditTask(taskDescription string) *editTask {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.SetValue(taskDescription)

	return &editTask{
		textInput: ti,
		task:      taskDescription,
	}
}

func (m editTask) Init() tea.Cmd {
	return textinput.Blink
}

func (m editTask) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		case "ctrl+c", "esc", "q":
			m.textInput.SetValue(m.task)
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m editTask) View() string {
	return m.textInput.View()
}

func EditTask(taskDescription string) string {
	p := tea.NewProgram(InitEditTask(taskDescription))
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return taskDescription
	}
	newTaskDescription := m.(editTask).textInput.Value()
	return strings.TrimSpace(newTaskDescription)
}
