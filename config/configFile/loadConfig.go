package configFile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-cli/config"
)

const CONFIG_FILE_NAME = "config.json"

func getGlobalConfigPath() string {
	return filepath.Join(config.GetGlobalAppDir(), CONFIG_FILE_NAME)
}

func LoadConfig() ConfigFile {
	configPath := getGlobalConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		defaultConfig := defaultConfig()
		configJSON, _ := json.MarshalIndent(defaultConfig, "", "  ")
		os.WriteFile(configPath, configJSON, 0644)
		fmt.Printf("⚙️  Config file created at\t(%s)\n", configPath)
		return defaultConfig
	}

	var config ConfigFile
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(configPath + " is corrupted: " + err.Error())
	}
	return config
}
