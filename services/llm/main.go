package llm

import (
	"sync"
	"todo-cli/config/configFile"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

type LlmService struct {
	url            string
	model          string
	apiKey         string
	clientRegistry *baml.ClientRegistry
	once           sync.Once
}

func InitLlmService(llmProvider configFile.LlmProvider) LlmService {
	return LlmService{
		url:    llmProvider.Url,
		model:  llmProvider.Model,
		apiKey: llmProvider.ApiKey,
	}
}
