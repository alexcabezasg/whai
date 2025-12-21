package commands

import (
	"fmt"
	"whai/pkg/utils/ui"
)

type WhaiCommand Command

func (c WhaiCommand) New() WhaiCommand {
	return WhaiCommand{
		Alias:       "whai",
		Description: "Reads the last error from the terminal and analyze it",
	}
}

func (c WhaiCommand) GetCommand() Command {
	return Command(c)
}

func (c WhaiCommand) AcceptsInput(input string) bool {
	return input == ""
}

func (c WhaiCommand) Run(args []string, ui ui.UI) error {
	fmt.Println("Whai Command")
	return nil
}
