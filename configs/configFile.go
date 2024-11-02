package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

var TODO_CLI_PATH = todoCliPath()
var CONFIG_PATH = TODO_CLI_PATH + "config.json"
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
	return homeDir + "/.todo-cli/"
}
