package agentsMd

import (
	"errors"
	"os"
)

const AGENTS_MD_FILE_NAME = "AGENTS.md"

func ReadAgentsMd() AgentsMd {
	content, err := os.ReadFile(AGENTS_MD_FILE_NAME)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return AgentsMd{Content: nil}
		}
		panic("Error when reading the agents.md file: " + err.Error())
	}
	str := string(content)
	return AgentsMd{Content: &str}
}
