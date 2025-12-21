package commands

import (
	"os"
	"os/exec"
	"whai/internal/config"
	"whai/internal/options"
	"whai/pkg/utils/logger"
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

func (c SetupCommand) Run(args []string, ui ui.UI, provider config.Provider) error {
	if len(args[1:]) == 0 {
		return OpenSettings()
	}

	return options.Run(args[1:], c.Options, provider, logger.NewLogger())
}

func OpenSettings() error {
	configPath, err := config.ConfigPath()
	if err != nil {
		return err
	}

	cmd := exec.Command("nano", configPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	logger.NewLogger().Info("Configuration successfully saved.")

	return nil
}
