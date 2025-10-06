package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) DeleteCompletedTasks() {
	t.db.Where("is_done = ?", true).Delete(&models.Task{})
}
