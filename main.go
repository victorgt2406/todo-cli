package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"todo-cli/commands"
	"todo-cli/configs/db"
	"todo-cli/controllers"
	"todo-cli/features"
	"todo-cli/models"
	"todo-cli/views/tdc"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	args := os.Args[1:]

	if noArgs(args) {
		p := tea.NewProgram(tdc.InitialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		if isCommand(args[0]) {
			commands.HandleCommand(args)
		} else {
			db := db.InitDB()
			taskController := controllers.NewTaskControllerWithDB(db)
			task := models.Task{Description: strings.Join(args, " ")}
			task.ID = taskController.CreateTask(task)
			err := features.SmartTask(db, &task)
			if err != nil {
				fmt.Println(err.Error())
			}
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
