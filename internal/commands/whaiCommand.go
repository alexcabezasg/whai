package commands

import (
	"fmt"
)

type WhaiCommand Command

func (c WhaiCommand) New() WhaiCommand {
	return WhaiCommand{
		Alias:       "whai",
		Description: "Reads the last error from the terminal and analyze it",
	}
}

func (c WhaiCommand) AcceptsInput(input string) bool {
	return input == ""
}

func (c WhaiCommand) Run(args []string) error {
	fmt.Println("Whai Command")
	return nil
}
