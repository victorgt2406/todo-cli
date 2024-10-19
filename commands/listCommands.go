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
		fmt.Println("[x] " + task.Description)
	} else {
		fmt.Println("[-] " + task.Description)
	}
}

func Edit(db *gorm.DB, task *models.Task) {
	originalDescription := task.Description
	description := views.EditTask(task.Description)
	db.Model(task).Update("Description", description)
	fmt.Println(originalDescription + " -> " + description)
}

func Delete(db *gorm.DB, task *models.Task) {
	db.Delete(task)
	fmt.Println("Task deleted!!")
}
