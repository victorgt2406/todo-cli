package llmService

import (
	"todo-cli/config/configFile"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

type LlmService struct {
	clientRegistry *baml.ClientRegistry
}

func InitLlmService(llmProvider configFile.LlmProvider) LlmService {
	clientRegistery := createCustomClientRegistry(llmProvider)
	return LlmService{
		clientRegistry: clientRegistery,
	}
}

func createCustomClientRegistry(llmProvider configFile.LlmProvider) *baml.ClientRegistry {

	clientRegistry := baml.NewClientRegistry()

	options := map[string]any{
		"base_url": llmProvider.Url,
		"model":    llmProvider.Model,
		"api_key":  llmProvider.ApiKey,
	}

	clientRegistry.AddLlmClient("customTodoCliProvider", "openai-generic", options)
	clientRegistry.SetPrimaryClient("customTodoCliProvider")

	return clientRegistry
}
