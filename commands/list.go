package commands

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"
	"todo-cli/views"
)

func List() {
	db := configs.InitDB()
	tasks := []models.Task{}
	db.Find(&tasks)

	if anyTasks(tasks) {
		fmt.Println("Your tasks:")

		selectedTask, err := views.SelectTask(tasks)

		if err != nil {
			return
		}

		fmt.Println("Selected task:", selectedTask.Description)
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
