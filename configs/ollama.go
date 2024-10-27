package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo-cli/models"
)

type Ollama struct {
	url string
}

func InitOllama() Ollama {
	config := LoadConfig()
	return Ollama{url: config.Ollama.Url}
}

func (o Ollama) Chat(context models.Context, message string) (string, error) {
	newMessage := models.ContextMessage{
		Role:    "user",
		Content: message,
	}
	url := o.url + "/api/chat"
	input := map[string]interface{}{
		"model":    context.Model,
		"messages": append(context.Messages, newMessage),
		"stream":   false,
	}

	// Convert input to JSON
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("error marshaling input: %w", err)
	}

	// POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonInput))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	// Check response is okay
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Read response
	var response struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		// Log the raw response body for debugging
		fmt.Printf("Raw response body: %s\n", string(body))
		return "", fmt.Errorf("failed to unmarshal Ollama response: %w\nRaw response: %s", err, string(body))
	}

	// Get message
	responseMessage := response.Message.Content

	return responseMessage, nil
}
