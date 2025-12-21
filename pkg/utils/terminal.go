package utils

import (
	"fmt"

	"github.com/pterm/pterm"
)

type CommandUI struct {
	Alias       string
	Description string
	Options     []OptionUI
}

type OptionUI struct {
	Flag        string
	Description string
	Values      []string
}

func PrintHelp(commands []CommandUI) {
	welcomeStyle := pterm.NewStyle(pterm.FgYellow, pterm.Bold)
	welcomeStyle.Println("Welcome to the help section of whai. Here are the commands: ")
	fmt.Println()
	for _, command := range commands {
		cmdTitle := pterm.NewStyle(pterm.FgCyan, pterm.Bold)
		cmdDescription := pterm.NewStyle(pterm.FgWhite)

		optionTitle := pterm.NewStyle(pterm.FgGreen, pterm.Bold)
		optionDescription := pterm.NewStyle(pterm.FgWhite)
		optionValues := pterm.NewStyle(pterm.FgYellow)

		cmdTitle.Println("whai " + command.Alias)
		cmdDescription.Println("Description: ", command.Description)
		cmdDescription.Println()

		for _, option := range command.Options {
			pterm.DefaultBox.WithTitle("Option").Println(optionTitle.Sprintf("%s", option.Flag))
			pterm.DefaultBox.WithTitle("Description").Println(optionDescription.Sprintf("%s", option.Description))

			if len(option.Values) > 0 {
				pterm.DefaultBox.WithTitle("Values").Println(optionValues.Sprintf("%v", option.Values))
			}
			fmt.Println()
		}
	}
}
