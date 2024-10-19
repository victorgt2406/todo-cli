package commands

var Commands = map[string]func(){
	"help": Help,
	"h":    Help,
	"l":    List,
	"li":   ListImportant,
	"r":    Remove,
	"ra":   RemoveAll,
}

func HandleCommand(args []string) {
	command := args[0]
	if _, commandExists := Commands[command]; commandExists {
		Commands[command]()
	} else {
		panic("Unknown command: " + command)
	}
}
