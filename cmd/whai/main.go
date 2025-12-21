package main

import (
	"fmt"
	"os"
	"whai/internal/commands"
	// "strings"
)

func main() {
	args := os.Args[1:]
	err := commands.Run(args, commands.GetAvailableCommands())
	if err != nil {
		fmt.Println(err)
		return
	}
}
