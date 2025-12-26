package commands

import (
	"os"
	"os/exec"
	"whai/internal/config"
	"whai/internal/options"
	"whai/pkg/context"
	"whai/pkg/utils/logger"
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
			options.RunnableOption(options.EditConfigOption{
				Flag:        "debug-mode",
				Description: "Enables the debug mode to see more traces.",
				Values:      []string{"true", "false"},
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

func (c SetupCommand) Run(args []string, ctx context.Context) error {
	if len(args[1:]) == 0 {
		return OpenSettings(ctx.Logger)
	}

	return options.Run(args[1:], c.Options, ctx)
}

func OpenSettings(logger logger.Logger) error {
	configPath, err := config.Path()
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

	logger.Info("Configuration successfully saved.")

	return nil
}
