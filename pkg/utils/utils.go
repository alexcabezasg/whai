package utils

import (
	"errors"
	"strings"
)

type Option struct {
	Key   string
	Value string
}

func ParseOption(arg string) (Option, error) {
	if !strings.Contains(arg, "--") {
		return Option{}, errors.New(arg + " is not a valid argument")
	}
	parsedArg := strings.Replace(arg, "--", "", 1)
	if strings.HasPrefix(parsedArg, "-") {
		return Option{}, errors.New(arg + " is not a valid argument")
	}

	splitStr := strings.Split(parsedArg, "=")
	if len(splitStr) != 2 {
		return Option{}, errors.New(arg + " is not a valid argument")
	}

	key := splitStr[0]
	value := splitStr[1]

	return Option{Key: key, Value: value}, nil
}

func ParseArguments(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}
