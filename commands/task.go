package commands

import (
	"fmt"
	"time"
	"todo-cli/configs"
	"todo-cli/models"
)

func CreateTask(description string) {
	db := configs.InitDB()
	task := models.Task{
		Description: description,
		Date:        time.Now().UTC().Format("2006-01-02T15:04:05"),
		CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05"),
		UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05"),
	}
	db.Create(&task)
	fmt.Println("Task created!!")
}
