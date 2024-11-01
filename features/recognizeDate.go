package features

import (
	"fmt"
	"time"
	"todo-cli/configs"
	"todo-cli/models"

	"gorm.io/gorm"
)

const DATE_CONTEXT_PATH = "./context/date_context.json"
const CATEGORY_CONTEXT_PATH = "./context/category_context.json"

func SetDateFromDescription(db *gorm.DB, task models.Task) error {
	ollama := configs.InitOllama()
	context, err := configs.LoadContext(DATE_CONTEXT_PATH)
	if err != nil {
		return fmt.Errorf("error loading date context")
	}
	message := createMessageForDate(task.Description)
	response, err := ollama.Chat(context, message)
	if err != nil {
		return fmt.Errorf("error when chatting with ollama")
	}
	if response != "INVALID" {
		date, err := time.Parse("2006-01-02", response)
		task.Date = &date
		if err != nil {
			return fmt.Errorf("error parsing date")
		}
		db.Model(&task).Where("id = ?", task.ID).Update("date", task.Date)
	}
	return nil
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
			"[In seven days] %s %s\n...",
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
