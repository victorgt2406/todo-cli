package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) DeleteTask(task models.Task) {
	t.db.Delete(&task)
}
