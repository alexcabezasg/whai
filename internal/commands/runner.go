package commands

import (
	"errors"
	"whai/pkg/utils"
)

func Run(args []string) error {
	input := utils.ParseArguments(args)
	for _, command := range GetAvailableCommands() {
		if command.AcceptsInput(input) {
			return command.Run(args)
		}
	}

	return errors.New("Command " + input + " not found. Try 'whai help' for more information.")
}

func GetAvailableCommands() []RunnableCommand {
	return []RunnableCommand{
		WhaiCommand{}.New(),
		HelpCommand{}.New(),
		SetupCommand{}.New(),
	}
}
