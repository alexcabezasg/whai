package commands

import (
	"whai/pkg/context"
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

func (c WhaiCommand) Run(_ []string, ctx context.Context) error {
	ctx.UI.RunWithSpinner("Analysing last command...", func() error {
		model := ctx.Config.Model
		ctx.Logger.Debug("Model selected: " + model)
		err, modelConfig := ctx.Config.GetModelConfiguration(model)
		if err != nil {
			ctx.Logger.Error("No model configuration found.")
			return err
		}
		ctx.Logger.Debug("Model configuration: url => " + modelConfig.URL + ", api_key => " + modelConfig.ApiKey)
		ctx.Logger.Debug("Model prompt: " + modelConfig.Prompt)

		err, aiResponse := ctx.AIResponseProvider.Get(modelConfig)
		if err != nil {
			ctx.Logger.Error("Error while getting ai response. " + err.Error())
			return err
		}

		if ctx.Config.OnlySuggest {
			ctx.UI.Println("Suggestion: "+aiResponse.Suggestion, ui.Format{Bold: true, Color: ui.Yellow})
			return nil
		}

		ctx.UI.Println("Summary: "+aiResponse.Summary, ui.Format{Color: ui.White})
		ctx.UI.Println("Root Cause: "+aiResponse.RootCause, ui.Format{Color: ui.Cyan})
		ctx.UI.Println("Suggestion: "+aiResponse.Suggestion, ui.Format{Bold: true, Color: ui.Yellow})

		return nil
	})

	return nil
}
