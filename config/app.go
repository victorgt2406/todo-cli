package config

import (
	"os"
	"path/filepath"
)

const APP_DIR = ".todo-cli"

func GetLocalAppDir() string {
	if err := os.MkdirAll(APP_DIR, 0755); err != nil {
		panic("Failed to create local app directory: " + err.Error())
	}
	return APP_DIR
}

func GetGlobalAppDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home directory: " + err.Error())
	}
	globalAppDir := filepath.Join(homeDir, APP_DIR)
	if err := os.MkdirAll(globalAppDir, 0755); err != nil {
		panic("Failed to create global app directory: " + err.Error())
	}
	return globalAppDir
}
