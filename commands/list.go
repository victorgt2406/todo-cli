package commands

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"
	"todo-cli/views"
)

func List() {
	fmt.Println("List")
	db := configs.InitDB()
	tasks := []models.Task{}
	db.Find(&tasks)
	var taskDescriptions []string
	for _, task := range tasks {
		taskDescriptions = append(taskDescriptions, task.Description)
	}
	selectedTask, err := views.SelectFromList(taskDescriptions)
	if err != nil {
		return
	}
	fmt.Println("Selected task:", selectedTask)
}

func ListImportant() {
	fmt.Println("ListImportant")
}
