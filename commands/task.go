package commands

import (
	"fmt"
	"time"
	"todo-cli/configs"
	"todo-cli/models"
)

const DATE_CONTEXT_PATH = "./context/date_context.json"
const CATEGORY_CONTEXT_PATH = "./context/category_context.json"

func CreateTask(description string) {
	db := configs.InitDB()
	config := configs.LoadConfig()
	task := models.Task{}
	task.Default()
	task.Description = description
	if config.Ollama.GetDateFromDescription {
		task.Date = getDateFromDescription(description)
	}
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

func getDateFromDescription(description string) string {
	context := configs.LoadContext(DATE_CONTEXT_PATH)
	ollama := configs.InitOllama()
	response, err := ollama.Chat(context, description)
	if err != nil {
		fmt.Println("Error getting date from description: ", err)
		fmt.Printf("\n⚠️ You can disable this feature in the config file.\n\n")
		return ""
	}
	fmt.Println(response)
	return response
}
