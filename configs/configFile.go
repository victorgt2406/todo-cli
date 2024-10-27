package configs

import (
	"encoding/json"
	"io"
	"os"
	"todo-cli/models"
)

type Config struct {
	Ollama struct {
		Url string `json:"url"`
	} `json:"ollama"`
	Features struct {
		RecognizeDate     bool `json:"recognizeDate"`
		RecognizeCategory bool `json:"recognizeCategory"`
	} `json:"features"`
}

func LoadConfig() Config {
	configJson, err := os.Open("config.json")
	if err != nil {
		panic("Error opening config.json: " + err.Error())
	}
	defer configJson.Close()

	byteValue, err := io.ReadAll(configJson)
	if err != nil {
		panic("Error reading config.json: " + err.Error())
	}

	var config Config
	json.Unmarshal(byteValue, &config)

	return config
}

func LoadContext(contextPath string) models.Context {
	contextFile, err := os.Open(contextPath)
	if err != nil {
		panic("Error opening context file: " + err.Error())
	}
	defer contextFile.Close()

	byteValue, err := io.ReadAll(contextFile)
	if err != nil {
		panic("Error reading context file: " + err.Error())
	}

	var context models.Context
	json.Unmarshal(byteValue, &context)
	return context
}
