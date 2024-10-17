package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	ID          int
	Description string
	Date        string
	CreatedAt   string
	UpdatedAt   string
}

var options = map[string]func(){
	"help": showHelp,
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		showHelp()
	} else {
		handleOption(args)
	}
}

func showHelp() {
	fmt.Println("Welcome to todo-cli (tdc)!")
	fmt.Println("\n🦙 Using LLMs it will set the date and task for you")
	fmt.Println("📂 Everything is stored locally in a `sqlite` database")
	fmt.Println("📅 If you want it can add them to your calendar")
	fmt.Println("\nUsage:")
	fmt.Println("  tdc make something for tomorrow")
}

func handleOption(args []string) {
	option := args[0]

	if _, optionExists := options[option]; optionExists {
		fmt.Println("Opción conocida:", option)
		options[option]()
	} else {
		taskDescription := strings.Join(args, " ")
		createTask(taskDescription)
	}
}

func initDB() *gorm.DB {
	dbPath := "./todo-cli.db"
	isNewDatabase := false
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		isNewDatabase = true
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Error opening DB: " + err.Error())
	}
	if isNewDatabase {
		fmt.Println("Setting up DB")
		err = db.AutoMigrate(&Task{})
		if err != nil {
			panic("Error migrating DB: " + err.Error())
		}
		fmt.Println("📂 DB created and migrated!")
	}
	return db
}

func createTask(description string) {
	db := initDB()
	task := Task{
		Description: description,
		Date:        time.Now().UTC().Format("2006-01-02T15:04:05"),
		CreatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05"),
		UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05"),
	}
	db.Create(&task)
	fmt.Println("Task created!!")
}
