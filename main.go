package main

import (
	"todo-cli/cli"
	"todo-cli/db"
)

func main() {
	db, context := db.InitDb()
	cli.Start(db, context)
}
