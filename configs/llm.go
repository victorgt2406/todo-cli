package configs

import (
	"todo-cli/configs/llm"
	"todo-cli/models"
)

type LLMClient struct {
	strategy llm.LLMStrategy
}

func NewLLMClient() *LLMClient {
	config := llm.Config{
		URL:    CONFIG.LLM.Url,
		APIKey: CONFIG.LLM.ApiKey,
		Model:  CONFIG.LLM.Model,
	}

	var strategy llm.LLMStrategy
	switch CONFIG.LLM.Provider {
	case "ollama":
		strategy = llm.NewOllamaStrategy(config)
	case "anthropic":
		strategy = llm.NewAnthropicStrategy(config)
	default:
		panic("Invalid LLM provider: " + CONFIG.LLM.Provider)
	}

	return &LLMClient{strategy: strategy}
}

func (llmClient *LLMClient) Chat(context models.Context, message string) (string, error) {
	return llmClient.strategy.Chat(context, message)
}
