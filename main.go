package main

import (
	"log"
	"os"
	"strings"
	"todo-cli/commands"
	"todo-cli/views/tdc"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	args := os.Args[1:]

	if noArgs(args) {
		// commands.List()
		p := tea.NewProgram(tdc.InitialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
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
