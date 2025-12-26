package commands

import (
	"strings"
	"whai/pkg/context"
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

func (c HelpCommand) Run(_ []string, ctx context.Context) error {
	ctx.UI.Println("Welcome to the help section of whai", ui.Format{Color: ui.Yellow, Bold: true})
	ctx.UI.EmptyLine()
	ctx.UI.Println("usage: whai [command]", ui.Format{Color: ui.Green, Bold: true})
	ctx.UI.EmptyLine()
	ctx.UI.Println("Available commands: ", ui.Format{Color: ui.Yellow, Bold: true})
	ctx.UI.EmptyLine()

	for _, command := range GetAvailableCommands() {
		command := command.GetCommand()

		ctx.UI.Println(command.Alias, ui.Format{Color: ui.Cyan, Bold: true})
		ctx.UI.Println("Description: "+command.Description, ui.Format{Bold: true})
		if len(command.Options) > 0 {
			ctx.UI.Println("Options: ", ui.Format{Bold: true})
			for _, option := range command.Options {
				option := option.GetOption()
				ctx.UI.Println("Flag: --"+option.Flag, ui.Format{Bold: true, Indentation: 2})
				ctx.UI.Println("Description: "+option.Description, ui.Format{Bold: true, Indentation: 2})
				ctx.UI.Print("Values: ", ui.Format{Bold: true, Indentation: 2})
				ctx.UI.Println(strings.Join(option.Values, ", "), ui.Format{Color: ui.Yellow, Bold: true})
				ctx.UI.EmptyLine()
			}
		}

		ctx.UI.EmptyLine()

	}

	return nil
}
