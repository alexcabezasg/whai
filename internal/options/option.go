package options

import (
	"slices"
	"whai/pkg/context"
	"whai/pkg/utils"
)

type Option struct {
	Flag        string
	Description string
	Values      []string
}

type RunnableOption interface {
	Run(arg string, ctx context.Context) error
	AcceptsInput(input string) bool
	GetOption() Option
}

func (opt Option) AcceptsInput(input string) bool {
	parsedOption, err := utils.ParseOption(input)
	if err != nil {
		return false
	}

	if !slices.Contains(opt.Values, parsedOption.Value) {
		return false
	}

	return parsedOption.Key == opt.Flag
}
