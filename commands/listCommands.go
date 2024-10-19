package commands

import (
	"fmt"
	"todo-cli/models"
	"todo-cli/views"

	"gorm.io/gorm"
)

func Enter(db *gorm.DB, task *models.Task) {
	db.Model(task).Update("IsDone", !task.IsDone)
	if task.IsDone {
		fmt.Println("✅ Task " + task.Description + " completed!!")
	} else {
		fmt.Println("❌ Task " + task.Description + " not completed!!")
	}
}

func Edit(db *gorm.DB, task *models.Task) {
	description := views.EditTask(task.Description)
	db.Model(task).Update("Description", description)
	fmt.Println("Task updated!!")
}

func Delete(db *gorm.DB, task *models.Task) {
	db.Delete(task)
	fmt.Println("Task deleted!!")
}
