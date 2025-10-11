package commands

func (c Command) getCommands() map[string]func(args []string) {
	return map[string]func(args []string){
		"--help":    c.help,
		"-h":        c.help,
		"--local":   c.createLocalDb,
		"--all":     c.readAllTasks,
		"-a":        c.readAllTasks,
		"--version": c.tdcVersion,
		"-v":        c.tdcVersion,
		"-c":        c.checkTask,
		"--check":   c.checkTask,
	}
}

func (c Command) handleRegisteredCommand(command string, args []string) bool {
	commands := c.getCommands()
	if handler, commandExists := commands[command]; commandExists {
		handler(args)
		return true
	}
	return false
}
