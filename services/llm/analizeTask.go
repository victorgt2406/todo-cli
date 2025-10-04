package llm

import (
	"context"
	"log"
	"time"

	b "todo-cli/baml_client"
	"todo-cli/models"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

func (LlmService) AnalizeTask(task models.Task) models.Task {
	return analizeTaskWithClient(task, "TogetherLlama3_3_70bTurbo")
}

func analizeTaskWithClient(task models.Task, clientName string) models.Task {
	ctx := context.Background()

	clientRegistry := baml.NewClientRegistry()
	clientRegistry.SetPrimaryClient(clientName)

	analizedTask, err := b.AnalizeTask(
		ctx,
		time.Now().Format("Monday 2006-01-02"),
		task.Description,
		b.WithClientRegistry(clientRegistry),
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
