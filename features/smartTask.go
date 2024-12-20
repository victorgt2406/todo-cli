package features

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"todo-cli/configs"
	"todo-cli/models"

	"gorm.io/gorm"
)

var CONTEXT_SMART_TASK_PATH = filepath.Join(configs.TODO_CLI_PATH, "context/smartTask.json")

func SmartTask(db *gorm.DB, task *models.Task) error {
	if !configs.CONFIG.Features.SmartTask {
		return nil
	}

	llm := configs.NewLLMClient()

	context, err := models.LoadContext(CONTEXT_SMART_TASK_PATH)
	if err != nil {
		return fmt.Errorf("[smartTask] error loading smart task context: %s", err.Error())
	}
	message := createMessage(task.Description)
	response, err := llm.Chat(context, message)
	if err != nil {
		return fmt.Errorf("[smartTask] error when chatting with llm: %s", err.Error())
	}
	description, date, err := validateResponse(response)
	if err != nil {
		return err
	}
	task.Description = description
	if date != nil {
		task.Date = date
	}
	db.Save(&task)
	return nil
}

func validateResponse(response string) (string, *time.Time, error) {
	split := strings.Split(response, "\n")
	if len(split) != 2 {
		return "", nil, fmt.Errorf("[smartTask] invalid response format: %s", response)
	}
	description, dateStr := split[0], split[1]
	if dateStr == "INVALID" {
		return description, nil, nil
	}
	date, err := time.Parse("2006-01-02T15:04", dateStr)
	if err != nil {
		return "", nil, fmt.Errorf("[smartTask] error parsing date: %s", response)
	}
	return description, &date, nil
}

func createMessage(description string) string {
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
		time.Now().Weekday(), time.Now().Format("2006-01-02T15:04"),
		time.Now().AddDate(0, 0, 1).Weekday(), time.Now().AddDate(0, 0, 1).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 2).Weekday(), time.Now().AddDate(0, 0, 2).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 3).Weekday(), time.Now().AddDate(0, 0, 3).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 4).Weekday(), time.Now().AddDate(0, 0, 4).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 5).Weekday(), time.Now().AddDate(0, 0, 5).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 6).Weekday(), time.Now().AddDate(0, 0, 6).Format("2006-01-02")+"T00:00",
		time.Now().AddDate(0, 0, 7).Weekday(), time.Now().AddDate(0, 0, 7).Format("2006-01-02")+"T00:00",
	)
}
