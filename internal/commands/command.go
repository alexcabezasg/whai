package commands

import (
	"whai/internal/options"
	"whai/pkg/context"
)

type Command struct {
	Alias       string
	Description string
	Options     []options.RunnableOption
}

type RunnableCommand interface {
	Run(args []string, ctx context.Context) error
	AcceptsInput(input string) bool
	GetCommand() Command
}

func (c Command) AcceptsInput(input string) bool {
	return input == c.Alias
}
