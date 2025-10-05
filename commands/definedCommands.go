package commands

var Commands = map[string]func(){
	"help":  help,
	"h":     help,
	"local": createLocalDb,
}

func handleRegisteredCommand(command string) bool {
	if _, commandExists := Commands[command]; commandExists {
		Commands[command]()
		return true
	}
	return false
}
