package commands

import (
	"errors"
	"whai/internal/config"
	"whai/pkg/utils"
	"whai/pkg/utils/ui"
)

func Run(args []string, commands []RunnableCommand) error {
	input := utils.ParseArguments(args)
	for _, command := range commands {
		if command.AcceptsInput(input) {
			return command.Run(args, ui.NewUI(), config.NewProvider())
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
