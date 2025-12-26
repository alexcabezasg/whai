package ai

import (
	"whai/internal/config"
)

type Result struct {
	Summary    string `json:"summary"`
	RootCause  string `json:"root_cause"`
	Suggestion string `json:"suggestion"`
}

type ResponseProvider interface {
	Get(configuration config.ModelConfiguration) (error, Result)
}

type DefaultResponseProvider struct{}

func NewAIResponseProvider() DefaultResponseProvider {
	return DefaultResponseProvider{}
}

func (p DefaultResponseProvider) Get(configuration config.ModelConfiguration) (error, Result) {
	// TODO Get the last failed command

	return nil, Result{
		Summary:    "Your last command failed due to this error: internal server error",
		RootCause:  "It might be possible that the server is not up and running.",
		Suggestion: "ping server.com to check the status of the server. If you don't get packets, contact support.",
	}
}
