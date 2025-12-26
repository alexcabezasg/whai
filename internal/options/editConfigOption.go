package options

import (
	"whai/pkg/context"
	"whai/pkg/utils"
)

type EditConfigOption Option

func (c EditConfigOption) GetOption() Option {
	return Option(c)
}

func (c EditConfigOption) AcceptsInput(input string) bool {
	return Option(c).AcceptsInput(input)
}

func (c EditConfigOption) Run(arg string, ctx context.Context) error {
	argument, _ := utils.ParseOption(arg)

	err := utils.SetFieldValue(&ctx.Config, argument)
	if err != nil {
		return err
	}

	return ctx.ConfigCommander.Set(ctx.Config)
}
