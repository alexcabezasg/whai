package commands

import (
	"whai/internal/options"
	"whai/pkg/utils"
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

func (c HelpCommand) Run(args []string) error {
	utils.PrintHelp(ToCommandUI(GetAvailableCommands()))
	return nil
}

func ToCommandUI(commands []RunnableCommand) []utils.CommandUI {
	var commandDTOs []utils.CommandUI
	for _, command := range commands {
		commandDTOs = append(commandDTOs, utils.CommandUI{
			Alias:       command.GetCommand().Alias,
			Description: command.GetCommand().Description,
			Options:     ToOptionUI(command.GetCommand().Options),
		})

	}

	return commandDTOs
}

func ToOptionUI(opts []options.RunnableOption) []utils.OptionUI {
	var optionDTOs []utils.OptionUI
	for _, option := range opts {
		optionDTOs = append(optionDTOs, utils.OptionUI{
			Flag:        option.GetOption().Flag,
			Description: option.GetOption().Description,
			Values:      option.GetOption().Values,
		})
	}

	return optionDTOs
}
