package llm

import "todo-cli/models"

// LLMStrategy defines the interface for different LLM providers
type LLMStrategy interface {
	Chat(context models.Context, message string) (string, error)
}

// Config holds common configuration for LLM providers
type Config struct {
	URL    string
	APIKey string
	Model  string
}
