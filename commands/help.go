package commands

import (
	"fmt"
	"todo-cli/config"
)

func help() {
	fmt.Println("Welcome to todo-cli (tdc)!")
	fmt.Println("\nA terminal-based todo list manager with AI to agile your workflow.")

	fmt.Println("\n━━━ COMMANDS ━━━")
	fmt.Println("  tdc local")
	fmt.Println("    Create a local database in current directory (.todo-cli/)")
	fmt.Println("    Use this for project-specific tasks")

	fmt.Println("\n━━━ USE SMART TASK FEATURE ━━━")
	fmt.Println("  Update the config file: " + config.GetGlobalAppDir() + "/config.json")
	fmt.Println("  Add the llm provider url, model and api key. And enable the smart task feature.")

	fmt.Println("\nFor more information, visit: https://github.com/victorgt2406/todo-cli")
}
