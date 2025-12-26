package ai

import (
	"whai/internal/config"
	"whai/internal/shell"
	"whai/pkg/utils/logger"
)

type Result struct {
	Summary    string `json:"summary"`
	RootCause  string `json:"root_cause"`
	Suggestion string `json:"suggestion"`
}

type ResponseProvider interface {
	Get(configuration config.ModelConfiguration, shellData shell.Data, logger logger.Logger) (error, Result)
}

func (r Result) String() string {
	return "{ summary: " + r.Summary + ", root_cause: " + r.RootCause + ", suggestion: " + r.Suggestion + " }"
}

func NewAIResponseProvider(cfg config.Config) ResponseProvider {
	switch cfg.Model {
	case "openai":
		return OpenAIResponseProvider{}
	default:
		panic("unsupported model: " + cfg.Model)
	}
}
