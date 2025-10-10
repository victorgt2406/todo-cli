package llmService

import (
	"todo-cli/config/agentsMd"

	baml "github.com/boundaryml/baml/engine/language_client_go/pkg"
)

type LlmService struct {
	clientRegistry *baml.ClientRegistry
	agentsMd       agentsMd.AgentsMd
}
