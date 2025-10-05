package commands

import (
	"fmt"
	"todo-cli/db"
)

func createLocalDb() {
	fmt.Println("Creating local database...")
	db.CreateDb(db.LocalDbPath())
	fmt.Println("\nLocal database created!")
}
