package ai

import (
	"context"
	"encoding/json"
	"whai/internal/shell"
	"whai/pkg/utils/logger"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"

	"whai/internal/config"
)

type OpenAIResponseProvider struct{}

func (OpenAIResponseProvider) Get(configuration config.ModelConfiguration, shellData shell.Data, logger logger.Logger) (error, Result) {
	client := openai.NewClient(
		option.WithAPIKey(configuration.ApiKey),
	)

	promptData := PromptData{Command: shellData.FailedCommand, Error: shellData.Error}

	logger.Debug("Loading openai.md prompt...")
	prompt, err := LoadPrompt("openai", promptData)
	if err != nil {
		logger.Error("Prompt for openai model could not be loaded")
		return err, Result{}
	}

	logger.Debug("Loading openai.md prompt complete")
	logger.Debug(prompt)

	logger.Debug("Calling openai configured model...")
	logger.Debug("Model: gpt-5-nano")

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5-nano",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
	})

	if err != nil {
		logger.Error("openai response could not be created")
		return err, Result{}
	}

	logger.Debug("openai model call success.")
	logger.Debug(resp.OutputText())

	var response Result
	err = json.Unmarshal([]byte(resp.OutputText()), &response)

	logger.Debug("Parsing response to appropriate struct...")

	if err != nil {
		return err, Result{}
	}

	logger.Debug("Parsing complete")
	logger.Debug(response.String())

	return nil, response
}
