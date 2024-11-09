package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var TODO_CLI_PATH = todoCliPath()
var CONFIG_PATH = filepath.Join(TODO_CLI_PATH, "config.json")
var CONFIG Config = LoadConfig()
var TODO_CLI_APP_NAME = "todo-cli"
var DB_PATH = filepath.Join(TODO_CLI_PATH, TODO_CLI_APP_NAME+".db")

type Config struct {
	Database struct {
		Provider string `json:"provider"`
		Url      string `json:"url"`
	} `json:"database"`
	Ollama struct {
		Url string `json:"url"`
	} `json:"ollama"`
	Features struct {
		SmartTask bool `json:"smartTask"`
	} `json:"features"`
}

func LoadConfig() Config {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		CreateConfig()
	}

	configJson, err := os.Open(CONFIG_PATH)
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

func CreateConfig() {
	err := os.MkdirAll(TODO_CLI_PATH, 0755)
	if err != nil {
		panic("Error creating config directory: " + err.Error())
	}

	defaultConfig := Config{}
	defaultConfig.Database.Provider = "sqlite"
	defaultConfig.Database.Url = DB_PATH
	defaultConfig.Ollama.Url = "http://localhost:11434"
	defaultConfig.Features.SmartTask = true

	jsonData, err := json.MarshalIndent(defaultConfig, "", "    ")
	if err != nil {
		panic("Error marshaling config: " + err.Error())
	}

	err = os.WriteFile(CONFIG_PATH, jsonData, 0644)
	if err != nil {
		panic("Error writing config file: " + err.Error())
	}
}

func todoCliPath() string {
	godotenv.Load()
	if os.Getenv("TDC_ENV") == "development" {
		fmt.Println("DEV_MODE")
		return "./"
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Error getting user home directory: " + err.Error())
	}
	return filepath.Join(homeDir, "."+TODO_CLI_APP_NAME)
}
