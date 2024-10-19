package main

import (
	"os"
	"strings"
	"todo-cli/commands"
)

func main() {
	args := os.Args[1:]

	if noArgs(args) {
		commands.Help()
	} else {
		if isCommand(args[0]) {
			commands.HandleCommand(args)
		} else {
			taskDescription := strings.Join(args, " ")
			commands.CreateTask(taskDescription)
		}
	}
}

func noArgs(args []string) bool {
	return len(args) == 0
}

func isCommand(arg1 string) bool {
	for command := range commands.Commands {
		if arg1 == command {
			return true
		}
	}
	return false
}
