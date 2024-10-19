package commands

import "fmt"

func Help() {
	fmt.Println("Welcome to todo-cli (tdc)!")
	fmt.Println("\n🦙 Using LLMs it will set the date and task for you")
	fmt.Println("📂 Everything is stored locally in a `sqlite` database")
	fmt.Println("📅 If you want it can add them to your calendar")
	fmt.Println("\nUsage:")
	fmt.Println("  tdc make something for tomorrow")
}
