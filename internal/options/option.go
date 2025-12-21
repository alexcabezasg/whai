package options

import (
	"slices"
	"whai/pkg/utils"
)

type Option struct {
	Flag        string
	Description string
	Values      []string
}

type RunnableOption interface {
	Run(args []string) error
	AcceptsInput(input string) bool
}

func (opt Option) AcceptsInput(input string) bool {
	parsedInput, err := utils.ParseArgument(input)
	if err != nil {
		return false
	}

	if !slices.Contains(opt.Values, parsedInput.Value) {
		return false
	}

	return parsedInput.Key == opt.Flag
}
