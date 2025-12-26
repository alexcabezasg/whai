package commands

import (
	"errors"
	"whai/pkg/context"
	"whai/pkg/utils"
)

func Run(args []string, ctx context.Context, commands ...RunnableCommand) error {
	input := utils.ParseArguments(args)

	if len(commands) == 0 {
		commands = GetAvailableCommands()
	}

	for _, command := range commands {
		if command.AcceptsInput(input) {
			return command.Run(args, ctx)
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
