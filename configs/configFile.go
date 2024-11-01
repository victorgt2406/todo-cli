package configs

import (
	"encoding/json"
	"io"
	"os"
)

var CONFIG Config = LoadConfig()

type Config struct {
	Ollama struct {
		Url string `json:"url"`
	} `json:"ollama"`
	Features struct {
		SmartTask bool `json:"smartTask"`
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
