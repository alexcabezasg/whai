package utils

import (
	"testing"
	"whai/internal/config"

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

func TestSetFieldValue(t *testing.T) {
	cfg := config.Config{Model: "gemini", OnlySuggest: true}

	err := SetFieldValue(&cfg, Option{Key: "only-suggest", Value: "false"})

	assert.Nil(t, err)
	assert.Equal(t, "gemini", cfg.Model)
	assert.Equal(t, false, cfg.OnlySuggest)

	err = SetFieldValue(&cfg, Option{Key: "model", Value: "chat-gpt"})

	assert.Nil(t, err)
	assert.Equal(t, "chat-gpt", cfg.Model)
	assert.Equal(t, false, cfg.OnlySuggest)

	err = SetFieldValue(&cfg, Option{Key: "key", Value: "value"})
	assert.Error(t, err)
}
