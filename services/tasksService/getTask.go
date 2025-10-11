package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	if result := t.db.First(&task, "id = ?", id); result.Error != nil {
		return task, result.Error
	}
	return task, nil
}
