package commands

import (
	"errors"
)

func Run(args []string) error {
	input := GetInputFromArgs(args)
	for _, command := range GetAvailableCommands() {
		if command.AcceptsInput(input) {
			return command.Run(args)
		}
	}

	return errors.New("Command " + input + " not found. Try 'whai help' for more information.")
}

func GetInputFromArgs(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

func GetAvailableCommands() []RunnableCommand {
	return []RunnableCommand{
		WhaiCommand{}.New(),
		HelpCommand{}.New(),
		SetupCommand{}.New(),
	}
}
