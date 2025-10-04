package main

import (
	"todo-cli/cli"
	"todo-cli/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, context := db.InitDb()
	cli.Start(db, context)
}
