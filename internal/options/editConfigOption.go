package options

import (
	"fmt"
	"whai/pkg/utils"
)

type EditConfigOption Option

func (c EditConfigOption) AcceptsInput(input string) bool {
	return Option(c).AcceptsInput(input)
}

func (c EditConfigOption) Run(args []string) error {
	argument, _ := utils.ParseOption(args[0])
	fmt.Println("setting ", argument.Key, " to ", argument.Value)

	return nil
}
