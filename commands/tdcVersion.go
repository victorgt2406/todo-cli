package commands

import "fmt"

var Version = "dev"

func (command Command) tdcVersion(args []string) {
	fmt.Println("tdc version:", Version)
}
