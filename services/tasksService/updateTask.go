package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) UpdateTask(task models.Task) {
	t.db.Save(&task)
}
