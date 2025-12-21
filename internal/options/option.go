package options

import "whai/pkg/utils"

type Option struct {
	Flag        string
	Description string
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

	return parsedInput.Key == opt.Flag
}
