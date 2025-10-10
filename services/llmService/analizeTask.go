package llmService

import (
	"context"
	"log"
	"time"

	b "todo-cli/baml_client"
	"todo-cli/models"
)

func (llmService LlmService) AnalizeTask(task models.Task) models.Task {
	ctx := context.Background()

	analizedTask, err := b.AnalizeTask(
		ctx,
		llmService.agentsMd.Content,
		time.Now().Format("Monday 2006-01-02"),
		task.Description,
		b.WithClientRegistry(llmService.clientRegistry),
	)
	if err != nil {
		log.Fatal(err)
	}
	task.Description = analizedTask.Description
	task.IsDone = analizedTask.IsDone
	if analizedTask.TodoDate != nil {
		task.TodoDate = convertStrTimeToTime(*analizedTask.TodoDate)
	}
	return task
}

func convertStrTimeToTime(strTime string) *time.Time {
	time, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		return nil
	}
	return &time
}
