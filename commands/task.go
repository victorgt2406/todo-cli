package commands

import (
	"fmt"
	"time"
	"todo-cli/configs"
	"todo-cli/models"
)

func CreateTask(description string) {
	db := configs.InitDB()
	task := models.Task{}
	task.Default()
	task.Description = description
	db.Create(&task)
	fmt.Println("Task created!!")
}

func UpdateTask(id int, description string) {
	db := configs.InitDB()
	task := models.Task{}
	db.First(&task, id)
	task.Description = description
	task.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	db.Save(&task)
	fmt.Println("Task updated!!")
}
