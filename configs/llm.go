package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"todo-cli/models"
)

type Llm struct {
	url    string
	apiKey string
	model  string
}

func InitLLM() Llm {
	switch CONFIG.LLM.Provider {
	case "ollama":
		return initOllama()
	case "anthropic":
		return initAnthropic()
	default:
		panic("Invalid LLM provider: " + CONFIG.LLM.Provider)
	}
}

func initOllama() Llm {
	return Llm{url: CONFIG.LLM.Url + "/api/chat", model: CONFIG.LLM.Model, apiKey: CONFIG.LLM.ApiKey}
}

func initAnthropic() Llm {
	return Llm{url: CONFIG.LLM.Url + "/messages", model: CONFIG.LLM.Model, apiKey: CONFIG.LLM.ApiKey}
}

func (llm Llm) Chat(context models.Context, message string) (string, error) {
	newMessage := models.ContextMessage{
		Role:    "user",
		Content: message,
	}
	url := llm.url
	input := map[string]interface{}{
		"model":    llm.model,
		"messages": append(context.Messages, newMessage),
		"stream":   false,
	}

	if CONFIG.LLM.Provider == "anthropic" {
		input["messages"].([]models.ContextMessage)[0].Role = "user"
		messagesSystemToUser(input["messages"].([]models.ContextMessage))
	}

	// Convert input to JSON
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return "", fmt.Errorf("error marshaling input: %w", err)
	}

	// Post
	body, err := post(url, llm.apiKey, jsonInput)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}

	// Read response
	var response struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}

	var anthropicResponse struct {
		Content []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"content"`
		Role string `json:"role"`
	}
	if CONFIG.LLM.Provider == "ollama" {
		err = json.Unmarshal(body, &response)
	} else if CONFIG.LLM.Provider == "anthropic" {
		err = json.Unmarshal(body, &anthropicResponse)
	}
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal Ollama response: %w\nRaw response: %s", err, string(body))
	}

	// Get message
	var responseMessage string
	if CONFIG.LLM.Provider == "ollama" {
		responseMessage = response.Message.Content
	} else if CONFIG.LLM.Provider == "anthropic" {
		responseMessage = anthropicResponse.Content[0].Text
	}
	return responseMessage, nil
}

func post(url string, apiKey string, body []byte) ([]byte, error) {
	if CONFIG.LLM.Provider == "anthropic" {
		body = append(body[:len(body)-1], []byte(`,"max_tokens":1024}`)...)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic("error creating request: " + err.Error())
	}
	if CONFIG.LLM.Provider == "anthropic" {
		req.Header.Set("x-api-key", apiKey)
		req.Header.Set("anthropic-version", "2023-06-01")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func messagesSystemToUser(messages []models.ContextMessage) {
	for _, message := range messages {
		if message.Role == "system" {
			message.Role = "user"
		}
	}
}
