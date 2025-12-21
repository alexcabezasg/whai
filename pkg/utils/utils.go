package utils

import (
	"errors"
	"strings"
)

type Argument struct {
	Key   string
	Value string
}

func ParseArgument(arg string) (Argument, error) {
	// --arg=value
	if !strings.Contains(arg, "--") {
		return Argument{}, errors.New(arg + " is not a valid argument")
	}
	parsedArg := strings.Replace(arg, "--", "", 1)
	if strings.HasPrefix(parsedArg, "-") {
		return Argument{}, errors.New(arg + " is not a valid argument")
	}

	splitStr := strings.Split(parsedArg, "=")
	if len(splitStr) != 2 {
		return Argument{}, errors.New(arg + " is not a valid argument")
	}

	key := splitStr[0]
	value := splitStr[1]

	return Argument{Key: key, Value: value}, nil
}
