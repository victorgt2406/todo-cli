package commands

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/models"
)

func List() {
	fmt.Println("List")
	db := configs.InitDB()
	tasks := []models.Task{}
	db.Find(&tasks)
	fmt.Println(tasks)
}

func ListImportant() {
	fmt.Println("ListImportant")
}
