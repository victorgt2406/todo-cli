package cli

import (
	"todo-cli/models"
	s "todo-cli/services/tasksService"
)

func (m model) getTasks() []models.Task {
	return m.tasksService.GetTasks(s.TaskFilter{}, s.OrderBy{})
}
