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
	default:
		panic("Invalid LLM provider: " + CONFIG.LLM.Provider)
	}
}

func initOllama() Llm {
	return Llm{url: CONFIG.LLM.Url, model: CONFIG.LLM.Model, apiKey: CONFIG.LLM.ApiKey}
}

func (llm Llm) Chat(context models.Context, message string) (string, error) {
	newMessage := models.ContextMessage{
		Role:    "user",
		Content: message,
	}
	url := llm.url + "/api/chat"
	input := map[string]interface{}{
		"model":    llm.model,
		"messages": append(context.Messages, newMessage),
		"stream":   false,
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

func post(url string, apiKey string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic("error creating request: " + err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
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
