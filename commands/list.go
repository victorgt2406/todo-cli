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

		var taskDescriptions []string
		for _, task := range tasks {
			taskDescriptions = append(taskDescriptions, task.Description)
		}

		selectedTask, err := views.SelectFromList(taskDescriptions)

		if err != nil {
			return
		}

		fmt.Println("Selected task:", selectedTask)
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
