package commands

import "whai/internal/options"

type SetupCommand Command

func (c SetupCommand) New() SetupCommand {
	return SetupCommand{
		Alias:       "setup",
		Description: "Adjust the configuration of whai",
		Options: []options.RunnableOption{
			options.RunnableOption(options.EditConfigOption{}.New(
				"only-suggest",
				"whai only prints the suggestion, not the reasoning. Default to false"),
			),
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

	return options.Run(args[1:], c.Options)
}

func OpenSettings() error {
	return nil
}
