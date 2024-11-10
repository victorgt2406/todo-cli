package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo-cli/models"
)

type OllamaStrategy struct {
	config Config
}

func NewOllamaStrategy(config Config) *OllamaStrategy {
	return &OllamaStrategy{
		config: Config{
			URL:    config.URL + "/api/chat",
			APIKey: config.APIKey,
			Model:  config.Model,
		},
	}
}

func (o *OllamaStrategy) Chat(context models.Context, message string) (string, error) {
	newMessage := models.ContextMessage{
		Role:    "user",
		Content: message,
	}

	input := map[string]interface{}{
		"model":    o.config.Model,
		"messages": append(context.Messages, newMessage),
		"stream":   false,
	}

	jsonInput, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("error marshaling input: %w", err)
	}

	body, err := o.post(jsonInput)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}

	var response struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal Ollama response: %w\nRaw response: %s", err, string(body))
	}

	return response.Message.Content, nil
}

func (o *OllamaStrategy) post(body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", o.config.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
