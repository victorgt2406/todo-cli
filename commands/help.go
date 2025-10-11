package commands

import (
	"fmt"
	"todo-cli/config"
)

func (command Command) help(args []string) {
	fmt.Println("Welcome to todo-cli (tdc)!")
	fmt.Println("\nA terminal-based todo list manager with AI to agile your workflow.")

	fmt.Println("\n━━━ COMMANDS ━━━")
	fmt.Println("  tdc")
	fmt.Println("    Launch interactive mode to manage tasks with a nice UX")
	fmt.Println("    This is the default when no arguments are provided")

	fmt.Println("\n  tdc [your task description]")
	fmt.Println("    Create a new task with the provided description")
	fmt.Println("    This is the default behavior when no command is specified")

	fmt.Println("\n  tdc --all, -a")
	fmt.Println("    Read all tasks")

	fmt.Println("\n  tdc --check, -c [task_id]")
	fmt.Println("    Toggle task status (done/pending) by its ID")

	fmt.Println("\n  tdc --local")
	fmt.Println("    Create a local database in current directory (.todo-cli/)")
	fmt.Println("    Use this for project-specific tasks")

	fmt.Println("\n  tdc --help, -h")
	fmt.Println("    Show this help message")

	fmt.Println("\n  tdc --version, -v")
	fmt.Println("    Show version information")

	fmt.Println("\n━━━ USE SMART TASK FEATURE ━━━")
	fmt.Println("  Update the config file: " + config.GetGlobalAppDir() + "/config.json")
	fmt.Println("  Add the llm provider url, model and api key. And enable the smart task feature.")

	fmt.Println("\nFor more information, visit: https://github.com/victorgt2406/todo-cli")
}
