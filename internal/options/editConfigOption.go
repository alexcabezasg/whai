package options

import (
	"fmt"
	"strconv"
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

	fmt.Println("DefaultModel: " + cfg.DefaultModel + ", OnlySuggest: " + strconv.FormatBool(cfg.OnlySuggest))

	return provider.Set(cfg)
}
