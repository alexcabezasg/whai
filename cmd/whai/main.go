package main

import (
	"os"
	"whai/internal/commands"
	"whai/pkg/utils/logger"
	// "strings"
)

func main() {
	args := os.Args[1:]
	err := commands.Run(args, commands.GetAvailableCommands())
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return
	}
}
