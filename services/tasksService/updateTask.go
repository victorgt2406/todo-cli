package tasksService

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

func UpdateTask(task models.Task, db *gorm.DB) {
	db.Save(&task)
}
