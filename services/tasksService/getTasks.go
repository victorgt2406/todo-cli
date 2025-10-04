package tasksService

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

func GetTasks(db *gorm.DB) []models.Task {
	var tasks []models.Task
	db.Find(&tasks)
	return tasks
}
