package commands

import (
	"whai/internal/options"
	"whai/pkg/utils/ui"
)

type SetupCommand Command

func (c SetupCommand) New() SetupCommand {
	return SetupCommand{
		Alias:       "setup",
		Description: "Adjust the configuration of whai. If no option provided, whai will open the configuration file.",
		Options: []options.RunnableOption{
			options.RunnableOption(options.EditConfigOption{
				Flag:        "only-suggest",
				Description: "whai only prints the suggestion, not the reasoning. Default to false",
				Values:      []string{"true", "false"},
			}),
			options.RunnableOption(options.EditConfigOption{
				Flag:        "default-model",
				Description: "Tells whai which model use on the next request.",
				Values:      []string{"chat-gpt", "gemini", "i don't know"},
			}),
		},
	}
}

func (c SetupCommand) GetCommand() Command {
	return Command(c)
}

func (c SetupCommand) AcceptsInput(input string) bool {
	return Command(c).AcceptsInput(input)
}

func (c SetupCommand) Run(args []string, ui ui.UI) error {
	if len(args[1:]) == 0 {
		return OpenSettings()
	}

	return options.Run(args[1:], c.Options)
}

func OpenSettings() error {
	return nil
}
