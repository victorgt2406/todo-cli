package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ContextMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Context struct {
	Model    string           `json:"model"`
	Messages []ContextMessage `json:"messages"`
}

func LoadContext(contextPath string) (Context, error) {
	contextFile, err := os.Open(contextPath)
	if err != nil {
		return Context{}, fmt.Errorf("error opening context file")
	}
	defer contextFile.Close()

	byteValue, err := io.ReadAll(contextFile)
	if err != nil {
		return Context{}, fmt.Errorf("error reading context file")
	}

	var context Context
	json.Unmarshal(byteValue, &context)
	return context, nil
}
