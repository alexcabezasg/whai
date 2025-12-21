package options

import (
	"whai/internal/config"
	"whai/pkg/utils"
)

type EditConfigOption Option

func (c EditConfigOption) GetOption() Option {
	return Option(c)
}

func (c EditConfigOption) AcceptsInput(input string) bool {
	return Option(c).AcceptsInput(input)
}

func (c EditConfigOption) Run(arg string, provider config.Provider) error {
	argument, _ := utils.ParseOption(arg)

	err, cfg := provider.Get()
	if err != nil {
		return err
	}

	err = utils.SetFieldValue(&cfg, argument)
	if err != nil {
		return err
	}

	return provider.Set(cfg)
}
