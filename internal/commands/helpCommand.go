package commands

import (
	"strings"
	"whai/internal/config"
	"whai/pkg/utils/ui"
)

type HelpCommand Command

func (c HelpCommand) New() HelpCommand {
	return HelpCommand{
		Alias:       "help",
		Description: "Shows this help",
	}
}

func (c HelpCommand) GetCommand() Command {
	return Command(c)
}

func (c HelpCommand) AcceptsInput(input string) bool {
	return Command(c).AcceptsInput(input)
}

func (c HelpCommand) Run(_ []string, userInterface ui.UI, _ config.Provider) error {
	userInterface.Println("Welcome to the help section of whai", ui.Format{Color: ui.Yellow, Bold: true})
	userInterface.EmptyLine()
	userInterface.Println("usage: whai [command]", ui.Format{Color: ui.Green, Bold: true})
	userInterface.EmptyLine()
	userInterface.Println("Available commands: ", ui.Format{Color: ui.Yellow, Bold: true})
	userInterface.EmptyLine()

	for _, command := range GetAvailableCommands() {
		command := command.GetCommand()

		userInterface.Println(command.Alias, ui.Format{Color: ui.Cyan, Bold: true})
		userInterface.Println("Description: "+command.Description, ui.Format{Bold: true})
		if len(command.Options) > 0 {
			userInterface.Println("Options: ", ui.Format{Bold: true})
			for _, option := range command.Options {
				option := option.GetOption()
				userInterface.Println("Flag: --"+option.Flag, ui.Format{Bold: true, Indentation: 2})
				userInterface.Println("Description: "+option.Description, ui.Format{Bold: true, Indentation: 2})
				userInterface.Print("Values: ", ui.Format{Bold: true, Indentation: 2})
				userInterface.Println(strings.Join(option.Values, ", "), ui.Format{Color: ui.Yellow, Bold: true})
				userInterface.EmptyLine()
			}
		}

		userInterface.EmptyLine()

	}

	return nil
}
