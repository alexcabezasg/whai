package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
	argument := ParseArguments([]string{"test"})
	assert.Equal(t, "test", argument)

	argument = ParseArguments([]string{})
	assert.Equal(t, "", argument)
}

func TestParseOption(t *testing.T) {
	opt, err := ParseOption("--key=value")
	assert.Nil(t, err)
	assert.Equal(t, Option{Key: "key", Value: "value"}, opt)

	_, err = ParseOption("key=value")
	assert.Error(t, err)

	_, err = ParseOption("---key=value")
	assert.Error(t, err)

	_, err = ParseOption("--key value")
	assert.Error(t, err)

	_, err = ParseOption("-key=value")
	assert.Error(t, err)
}
