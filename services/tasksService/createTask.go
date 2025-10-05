package tasksService

import (
	"todo-cli/models"
)

func (t TasksService) CreateTask(description string) models.Task {
	newTask := models.Task{
		Description: description,
	}
	t.db.Create(&newTask)
	return newTask
}
