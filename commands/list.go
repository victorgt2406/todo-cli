package commands

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"
	"todo-cli/views"

	"gorm.io/gorm"
)

var ListCommands = map[string]func(db *gorm.DB, task *models.Task){
	"enter": Enter,
	"e":     Edit,
	"d":     Delete,
}

func List() {
	db := configs.InitDB()
	tasks := []models.Task{}
	db.Find(&tasks)

	if anyTasks(tasks) {
		selectedTask, command, err := views.SelectTask(tasks)
		if err != nil {
			return
		}
		ListCommands[command](db, selectedTask)

	} else {
		fmt.Println("No tasks found...")
	}

}

func ListImportant() {
	fmt.Println("ListImportant")
}

func anyTasks(tasks []models.Task) bool {
	return len(tasks) > 0
}
