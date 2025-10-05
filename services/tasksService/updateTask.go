package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) UpdateTask(task models.Task) models.Task {
	t.db.Save(task)
	return task
}
