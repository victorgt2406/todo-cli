package tdc

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/features"
	"todo-cli/models"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

var colors = map[string]lipgloss.Color{
	"red":    lipgloss.Color("1"),
	"green":  lipgloss.Color("2"),
	"yellow": lipgloss.Color("3"),
	"blue":   lipgloss.Color("4"),
	"purple": lipgloss.Color("5"),
	"dark":   lipgloss.Color("#1c1c1c"),
	"light":  lipgloss.Color("#f0f0f0"),
}

var styles = map[string]lipgloss.Style{
	"selectedNotChecked": lipgloss.NewStyle().Foreground(colors["yellow"]).Bold(true),
	"selectedChecked":    lipgloss.NewStyle().Foreground(colors["yellow"]).Faint(true),
	"checked":            lipgloss.NewStyle().Foreground(colors["dark"]).Faint(true),
	"notChecked":         lipgloss.NewStyle().Foreground(colors["light"]),
	"footer":             lipgloss.NewStyle().Foreground(colors["blue"]).Italic(true),
	"title":              lipgloss.NewStyle().Foreground(colors["light"]).Bold(true).Padding(0, 1, 0, 1).Align(lipgloss.Left).Background(colors["blue"]),
}

type contextState string

const (
	contextEditTask contextState = "editTask"
	contextNewTask  contextState = "newTask"
	contextTasks    contextState = "tasks"
)

type model struct {
	db        *gorm.DB
	tasks     []models.Task
	cursor    int
	context   contextState
	textInput textinput.Model
}

type TasksUpdatedMsg struct{}

func InitialModel() model {
	db := configs.InitDB()
	var tasks []models.Task
	db.Find(&tasks)
	ti := textinput.New()
	ti.Prompt = ""
	return model{
		db:        db,
		tasks:     tasks,
		cursor:    0,
		context:   contextTasks,
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.context {
	case contextTasks:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q", "Q", "esc":
				return m, tea.Quit
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}
			case " ":
				Check(m)
			case "enter":
				Check(m)
				return m, tea.Quit
			case "e", "E":
				if len(m.tasks) > 0 {
					m.context = contextEditTask
					m.textInput.Focus()
					m.textInput.SetValue(m.tasks[m.cursor].Description)
				}
			case "n", "N":
				m.context = contextNewTask
				m.textInput.Focus()
				m.textInput.Placeholder = "Create a new task"
			case "delete", "d", "backspace":
				if m.cursor >= 0 && m.cursor < len(m.tasks) {
					m.db.Delete(&m.tasks[m.cursor])
					m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
				}
			}
		case TasksUpdatedMsg:
			m.db.Find(&m.tasks)
		}
	case contextNewTask, contextEditTask:
		return textInputUpdate(m, msg)
	}
	return m, nil
}

func textInputUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.context == contextNewTask {
				m.cursor = len(m.tasks)
				m.tasks = append(m.tasks, models.Task{})
			}
			newDescription := m.textInput.Value()
			m.tasks[m.cursor].Description = newDescription
			m.db.Save(&m.tasks[m.cursor])
			m.textInput.Reset()
			m.textInput.Blur()
			m.context = contextTasks
			return m, tea.Batch(
				func() tea.Msg { return nil },
				func() tea.Msg {
					err := features.SetDateFromDescription(m.db, m.tasks[m.cursor])
					if err != nil {
						fmt.Println(err)
					}
					m.db.Find(&m.tasks)
					return TasksUpdatedMsg{}
				},
			)
		case "esc":
			m.textInput.Reset()
			m.textInput.Blur()
			m.context = contextTasks
		}
	}
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// The header
	title := styles["title"].Render("Todo cli")
	s := fmt.Sprintf("%s\n", title)

	// Body: Tasks
	for i, task := range m.tasks {
		cursor := " "
		if m.cursor == i && m.context != contextNewTask {
			cursor = ">"
		}
		checked := " "
		if task.IsDone {
			checked = "x"
		}
		dateStr := ""
		if task.Date != nil {
			dateStr = fmt.Sprintf(" (%s)", task.Date.Format("Mon 02/01/2006"))
		}
		body := fmt.Sprintf("%s%s", task.Description, dateStr)
		if m.cursor == i && m.context == contextEditTask {
			body = m.textInput.View()
		}
		taskElement := fmt.Sprintf("%s [%s] %s", cursor, checked, body)
		if m.cursor == i && m.context != contextNewTask {
			if task.IsDone {
				taskElement = styles["selectedChecked"].Render(taskElement)
			} else {
				taskElement = styles["selectedNotChecked"].Render(taskElement)
			}
		} else if task.IsDone {
			taskElement = styles["checked"].Render(taskElement)
		} else {
			taskElement = styles["notChecked"].Render(taskElement)
		}
		s += taskElement + "\n"
	}
	if m.context == contextNewTask {
		m.textInput.Prompt = "> [ ] "
		m.textInput.PromptStyle = styles["selectedNotChecked"]
		m.textInput.TextStyle = styles["selectedNotChecked"]
		s += m.textInput.View()
	}
	// Footer
	footerMessage := "Press \"q\" to quit, \"e\" to edit, \"n\" to create, \"d\"/\"backspace\"/\"delete\" to delete"
	if m.context == contextEditTask || m.context == contextNewTask {
		footerMessage = "Press esc to cancel, enter to save"
	}
	footer := styles["footer"].Render(footerMessage)
	s += fmt.Sprintf("\n%s\n", footer)
	return s
}
