package commands

import (
	"fmt"
	"todo-cli/db"
)

func (command Command) createLocalDb(args []string) {
	fmt.Println("Creating local database...")
	db.CreateDb(db.LocalDbPath())
	fmt.Println("\nLocal database created!")
}
