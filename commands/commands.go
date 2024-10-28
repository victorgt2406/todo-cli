package commands

import "fmt"

var Commands = map[string]func(){
	"help": Help,
	"h":    Help,
}

func HandleCommand(args []string) {
	command := args[0]
	if _, commandExists := Commands[command]; commandExists {
		Commands[command]()
	} else {
		fmt.Println("Unknown command: " + command)
	}
}
