package commands

import (
	"fmt"
	"todo-cli/models"
	"todo-cli/views/viewTask"

	"gorm.io/gorm"
)

var ListCommands = map[string]func(db *gorm.DB, task *models.Task){
	"enter": toggleTaskStatus,    //Enter
	" ":     toggleTaskStatus,    //Enter
	"e":     editTaskDescription, //Edit
	"d":     deleteTask,          //Delete
	"del":   deleteTask,          //Delete
}

func toggleTaskStatus(db *gorm.DB, task *models.Task) {
	db.Model(task).Update("IsDone", !task.IsDone)
	if task.IsDone {
		fmt.Println("[x] " + task.Description)
	} else {
		fmt.Println("[-] " + task.Description)
	}
}

func editTaskDescription(db *gorm.DB, task *models.Task) {
	originalDescription := task.Description
	description := viewTask.Task(task.Description)
	db.Model(task).Update("Description", description)
	fmt.Println(originalDescription + " -> " + description)
}

func deleteTask(db *gorm.DB, task *models.Task) {
	// TODO: Add confirmation
	db.Delete(task)
	fmt.Println("Task deleted")
}
