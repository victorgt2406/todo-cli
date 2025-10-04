package tasksService

import (
	"todo-cli/models"

	"gorm.io/gorm"
)

func DeleteTask(task models.Task, db *gorm.DB) {
	db.Delete(&task)
}
