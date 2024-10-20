package viewTask

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func Task(taskDescription string) string {
	p := tea.NewProgram(initTask(taskDescription))
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return taskDescription
	}
	newTaskDescription := m.(teaModel).textInput.Value()
	return strings.TrimSpace(newTaskDescription)
}

type teaModel struct {
	textInput textinput.Model
	task      string
}

func initTask(taskDescription string) *teaModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.SetValue(taskDescription)

	return &teaModel{
		textInput: ti,
		task:      taskDescription,
	}
}

func (m teaModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m teaModel) View() string {
	return m.textInput.View()
}
