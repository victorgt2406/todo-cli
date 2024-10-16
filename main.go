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
	ID        int
	Task      string
	Date      string
	CreatedAt string
	UpdatedAt string
}

var options = map[string]func(){
	"help":     showHelp,
	"llm":      setupLLM,
	"db":       checkDB,
	"calendar": syncCalendar,
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
		taskDefinition := strings.Join(args, " ")
		createTask(taskDefinition)
	}
}

func setupLLM() {
	fmt.Println("Setting up LLM")
}

func checkDB() {
	if _, err := os.Stat("./todo-cli.db"); os.IsNotExist(err) {
		fmt.Println("Setting up DB")
		db, err := gorm.Open(sqlite.Open("./todo-cli.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("Error creating DB:", err)
			panic("Error creating DB")
		}
		db.AutoMigrate(&Task{})
		fmt.Println("DB created")
	} else {
		fmt.Println("DB already exists, skipping setup")
	}

}

func connectDb() *gorm.DB {
	checkDB()
	db, err := gorm.Open(sqlite.Open("./todo-cli.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening DB:", err)
		panic("Error opening DB")
	}
	fmt.Println("Connected to DB")
	return db
}

func syncCalendar() {
	fmt.Println("Syncing calendar")
}

func createTask(taskDefinition string) {
	fmt.Println("Creating task:", taskDefinition)
	db := connectDb()
	task := Task{
		Task:      taskDefinition,
		Date:      time.Now().UTC().Format("2006-01-02T15:04:05"),
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05"),
		UpdatedAt: time.Now().UTC().Format("2006-01-02T15:04:05"),
	}
	db.Create(&task)
	fmt.Println("Task created!!")
}
