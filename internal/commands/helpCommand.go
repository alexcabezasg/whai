package commands

import "fmt"

type HelpCommand Command

func (c HelpCommand) New() HelpCommand {
	return HelpCommand{
		Alias:       "help",
		Description: "Shows this help",
	}
}

func (c HelpCommand) AcceptsInput(input string) bool {
	return Command(c).AcceptsInput(input)
}

func (c HelpCommand) Run(args []string) error {
	fmt.Println("Help Command")
	return nil
}
