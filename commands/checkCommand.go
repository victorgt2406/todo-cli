package commands

import (
	"os"
	"strings"
)

func (command Command) IsCommand() bool {
	args := getArgs()
	if anyArgs(args) {
		if !command.handleRegisteredCommand(args[0], args[1:]) {

			description := strings.Join(args, " ")
			task := command.tasksService.CreateTask(description)

			if command.features.SmartTask {
				task = command.llmService.AnalizeTask(task)
				command.tasksService.UpdateTask(task)
			}
		}
		return true
	}

	return false
}

func getArgs() []string {
	return os.Args[1:]
}

func anyArgs(args []string) bool {
	return len(args) > 0
}
