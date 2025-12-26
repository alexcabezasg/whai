package ai

import (
	"bytes"
	"os"
	"text/template"
)

type PromptData struct {
	Command string
	Error   string
}

func LoadPrompt(model string, data PromptData) (string, error) {
	tplBytes, err := os.ReadFile("internal/prompts/" + model + ".md")
	if err != nil {
		return "", err
	}

	tpl, err := template.New("prompt").Parse(string(tplBytes))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
