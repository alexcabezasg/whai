package commands

import (
	"errors"
	"fmt"
	"whai/pkg/utils"
)

type SetupCommand Command

func (c SetupCommand) New() SetupCommand {
	return SetupCommand{
		Alias:       "setup",
		Description: "Adjust the configuration of whai",
		SubCommands: []Command{
			{
				Alias:       "only-suggest",
				Description: "whai only prints the suggestion, not the reasoning. Default to false",
			},
		},
	}
}

func (c SetupCommand) AcceptsInput(input string) bool {
	return Command(c).AcceptsInput(input)
}

func (c SetupCommand) Run(args []string) error {
	if len(args[1:]) == 0 {
		return OpenSettings()
	}

	return c.RunSubCommands(args[1:])
}

func OpenSettings() error {
	return nil
}

func (c SetupCommand) RunSubCommands(args []string) error {
	var errs []error
	for _, arg := range args {
		IsSupported := false
		for _, subCommand := range c.SubCommands {
			parsedArgument, err := utils.ParseArgument(arg)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			if subCommand.AcceptsInput(parsedArgument.Key) {
				// Get the configuration file
				// Set new value to key
				IsSupported = true
				fmt.Println("setting ", parsedArgument.Key, " to ", parsedArgument.Value)
				continue
			}
		}
		if !IsSupported {
			errs = append(errs, errors.New("Argument "+arg+" not supported."))
		}
	}

	return errors.Join(errs...)
}
