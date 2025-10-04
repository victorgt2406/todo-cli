package tasksService

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

func CreateTask(description string, db *gorm.DB) models.Task {
	newTask := models.Task{
		Description: description,
	}
	db.Create(&newTask)
	return newTask
}
