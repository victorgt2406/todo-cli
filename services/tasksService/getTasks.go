package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) GetTasks() []models.Task {
	var tasks []models.Task
	t.db.Find(&tasks)
	return tasks
}
