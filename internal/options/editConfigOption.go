package options

import (
	"errors"
	"fmt"
	"reflect"
	"whai/internal/config"
	"whai/pkg/utils"

	"github.com/iancoleman/strcase"
)

type EditConfigOption Option

func (c EditConfigOption) GetOption() Option {
	return Option(c)
}

func (c EditConfigOption) AcceptsInput(input string) bool {
	return Option(c).AcceptsInput(input)
}

func (c EditConfigOption) Run(args []string, provider config.Provider) error {
	argument, _ := utils.ParseOption(args[0])

	err, cfg := provider.Get()
	if err != nil {
		return err
	}

	configKey := strcase.ToCamel(argument.Key)
	field := reflect.ValueOf(&cfg).Elem().FieldByName(configKey)

	switch field.Kind() {
	case reflect.Bool:
		field.SetBool(argument.Value == "true")

	case reflect.String:
		field.SetString(argument.Value)
	default:
		return errors.New("unexpected error while inferring the configuration change")
	}

	fmt.Println("setting ", argument.Key, " to ", argument.Value)

	return provider.Set(cfg)
}
