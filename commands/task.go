package commands

import (
	"fmt"
	"sync"
	"time"
	"todo-cli/configs"
	"todo-cli/models"
	"todo-cli/views/load"

	"gorm.io/gorm"
)

const DATE_CONTEXT_PATH = "./context/date_context.json"
const CATEGORY_CONTEXT_PATH = "./context/category_context.json"

func CreateTask(description string) {
	db := configs.InitDB()
	task := models.Task{}
	task.Default()
	task.Description = description
	result := db.Create(&task)
	if result.Error != nil {
		fmt.Println("Error creating task:", result.Error)
		return
	}
	fmt.Printf("Task saved #%d\n", task.ID)
	ProcessTask(db, task)
}

func ProcessTask(db *gorm.DB, task models.Task) {

	config := configs.LoadConfig()
	if config.Features.RecognizeDate {
		var wg sync.WaitGroup
		wg.Add(1)

		// get date from description
		loader := load.Load([]string{
			"Getting date from description",
		})
		go func() {
			defer wg.Done()

			getDateFromDescription(db, task)
			loader("Getting date from description")
		}()

		wg.Wait()
	}
}

func UpdateTask(id int, description string) {
	db := configs.InitDB()
	task := models.Task{}
	db.First(&task, id)
	task.Description = description
	now := time.Now().UTC()
	task.UpdatedAt = &now
	db.Save(&task)
	fmt.Println("Task updated!!")
}

func getDateFromDescription(db *gorm.DB, task models.Task) {
	context := configs.LoadContext(DATE_CONTEXT_PATH)
	ollama := configs.InitOllama()
	message := createMessageForDate(task.Description)
	response, err := ollama.Chat(context, message)
	if err != nil {
		fmt.Println("Error getting date from description: ", err)
	}
	if response != "INVALID" {
		*task.Date, err = time.Parse("2006-01-02", response)
		if err != nil {
			fmt.Println("Error parsing date: ", err)
		}
		db.Model(&task).Where("id = ?", task.ID).Update("date", task.Date)
	}
}

func createMessageForDate(description string) string {
	return fmt.Sprintf(
		"[Task Description] %s\n"+
			"[Current Date - Today] %s %s\n"+
			"[In one day - Tomorrow] %s %s\n"+
			"[In two days - Day after tomorrow] %s %s\n"+
			"[In three days] %s %s\n"+
			"[In four days] %s %s\n"+
			"[In five days] %s %s\n"+
			"[In six days] %s %s\n"+
			"[In seven days] %s %s",
		description,
		time.Now().Weekday(), time.Now().Format("2006-01-02"),
		time.Now().AddDate(0, 0, 1).Weekday(), time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 2).Weekday(), time.Now().AddDate(0, 0, 2).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 3).Weekday(), time.Now().AddDate(0, 0, 3).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 4).Weekday(), time.Now().AddDate(0, 0, 4).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 5).Weekday(), time.Now().AddDate(0, 0, 5).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 6).Weekday(), time.Now().AddDate(0, 0, 6).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 7).Weekday(), time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	)
}
