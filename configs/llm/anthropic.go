package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo-cli/models"
)

type AnthropicStrategy struct {
	config Config
}

func NewAnthropicStrategy(config Config) *AnthropicStrategy {
	return &AnthropicStrategy{
		config: Config{
			URL:    config.URL + "/messages",
			APIKey: config.APIKey,
			Model:  config.Model,
		},
	}
}

func (a *AnthropicStrategy) Chat(context models.Context, message string) (string, error) {
	newMessage := models.ContextMessage{
		Role:    "user",
		Content: message,
	}

	messages := append(context.Messages, newMessage)
	messages[0].Role = "user"
	a.messagesSystemToUser(messages)

	input := map[string]interface{}{
		"model":    a.config.Model,
		"messages": messages,
		"stream":   false,
	}

	jsonInput, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("error marshaling input: %w", err)
	}

	body, err := a.post(jsonInput)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}

	var response struct {
		Content []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"content"`
		Role string `json:"role"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal Anthropic response: %w\nRaw response: %s", err, string(body))
	}

	return response.Content[0].Text, nil
}

func (a *AnthropicStrategy) post(body []byte) ([]byte, error) {
	body = append(body[:len(body)-1], []byte(`,"max_tokens":1024}`)...)

	req, err := http.NewRequest("POST", a.config.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", a.config.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

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

func (a *AnthropicStrategy) messagesSystemToUser(messages []models.ContextMessage) {
	for _, message := range messages {
		if message.Role == "system" {
			message.Role = "user"
		}
	}
}
