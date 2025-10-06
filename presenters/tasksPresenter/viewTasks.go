package tasksPresenter

import (
	"fmt"
	"time"
	"todo-cli/models"
	"todo-cli/utils"

	"github.com/charmbracelet/bubbles/textinput"
)

type RenderProps struct {
	Tasks       []models.Task
	Cursor      int
	ViewContext models.ViewContext
	TextInput   textinput.Model
}

func (t TasksPresenter) Render(p RenderProps) string {
	// The header
	s := fmt.Sprintf("%s\n\n", t.title())

	// Tasks list
	isCompletedTasks := false
	for i, task := range p.Tasks {
		description := task.Description
		if p.ViewContext == models.ViewEditTask && p.Cursor == i {
			description = p.TextInput.View()
		}
		if task.IsDone && !isCompletedTasks {
			if p.ViewContext == models.ViewNewTask {
				p.TextInput.Prompt = "> [ ] "
				s += p.TextInput.View() + "\n"
			}
			s += "\nCompleted Tasks:\n"
			isCompletedTasks = true
		}
		s += t.task(taskProps{
			description: description,
			todoDate:    task.TodoDate,
			isDone:      task.IsDone,
			selected:    p.Cursor == i && p.ViewContext != models.ViewNewTask,
			viewContext: p.ViewContext,
		})
		s += "\n"
	}

	// The footer
	s += fmt.Sprintf("\n%s\n", t.footer())

	// Send the UI for rendering
	return s
}

type taskProps struct {
	description string
	todoDate    *time.Time
	isDone      bool
	selected    bool
	viewContext models.ViewContext
}

func (t TasksPresenter) task(p taskProps) string {
	s := ""

	strCursor := " "
	checked := " "
	description := p.description
	if p.selected {
		strCursor = ">"
	}

	if p.isDone {
		checked = "x"
	}

	s += fmt.Sprintf("%s [%s] %s", strCursor, checked, description)
	if p.todoDate != nil {
		s += fmt.Sprintf(" 📅 %s", utils.FormatDateToString(*p.todoDate))
	}
	if p.selected {
		if p.isDone {
			s = utils.Styles["selectedChecked"].Render(s)
		} else {
			s = utils.Styles["selectedNotChecked"].Render(s)
		}
	} else {
		if p.isDone {
			s = utils.Styles["checked"].Render(s)
		} else {
			s = utils.Styles["notChecked"].Render(s)
		}
	}
	s = utils.Styles["task"].Render(s)
	return s
}

func (t TasksPresenter) footer() string {
	return utils.Styles["footer"].Render("q: quit | n: new task | e: edit | <space>: toggle completion")
}

func (t TasksPresenter) title() string {
	return utils.Styles["title"].Render("Todo List")
}
