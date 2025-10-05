package llm

import (
	"context"
	"log"
	"time"

	b "todo-cli/baml_client"
	"todo-cli/models"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

func (llmService *LlmService) AnalizeTask(task models.Task) models.Task {
	clientRegistry := llmService.getCustomClientRegistry()
	return analizeTaskWithClient(task, clientRegistry)
}

func (llmService *LlmService) getCustomClientRegistry() *baml.ClientRegistry {
	llmService.once.Do(func() {
		clientRegistry := baml.NewClientRegistry()

		options := map[string]any{
			"base_url": llmService.url,
			"model":    llmService.model,
			"api_key":  llmService.apiKey,
		}

		clientRegistry.AddLlmClient("customTodoCliProvider", "openai-generic", options)
		clientRegistry.SetPrimaryClient("customTodoCliProvider")

		llmService.clientRegistry = clientRegistry
	})

	return llmService.clientRegistry
}

func analizeTaskWithClient(task models.Task, clientRegistry *baml.ClientRegistry) models.Task {
	ctx := context.Background()

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
