package commands

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"
	"todo-cli/views/viewTasks"
)

func List() {
	db := configs.InitDB()
	tasks := []models.Task{}
	db.Find(&tasks)

	if anyTasks(tasks) {
		allowedCommands := getAllowedCommands()
		selectedTask, command, err := viewTasks.Tasks(tasks, allowedCommands)
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

func getAllowedCommands() []string {
	allowedCommands := make([]string, 0, len(ListCommands))
	for command := range ListCommands {
		allowedCommands = append(allowedCommands, command)
	}
	return allowedCommands
}
