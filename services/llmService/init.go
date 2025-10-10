package llmService

import (
	"todo-cli/config/agentsMd"
	"todo-cli/config/configFile"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

type InitLlmServiceProps struct {
	LlmProvider configFile.LlmProvider
	AgentsMd    agentsMd.AgentsMd
}

func InitLlmService(p InitLlmServiceProps) LlmService {
	clientRegistery := createCustomClientRegistry(p.LlmProvider)
	return LlmService{
		clientRegistry: clientRegistery,
		agentsMd:       p.AgentsMd,
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
