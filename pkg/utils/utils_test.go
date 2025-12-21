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
	cfg := config.Config{DefaultModel: "gemini", OnlySuggest: true}

	err := SetFieldValue(&cfg, Option{Key: "only-suggest", Value: "false"})

	assert.Nil(t, err)
	assert.Equal(t, "gemini", cfg.DefaultModel)
	assert.Equal(t, false, cfg.OnlySuggest)

	err = SetFieldValue(&cfg, Option{Key: "default-model", Value: "chat-gpt"})

	assert.Nil(t, err)
	assert.Equal(t, "chat-gpt", cfg.DefaultModel)
	assert.Equal(t, false, cfg.OnlySuggest)

	err = SetFieldValue(&cfg, Option{Key: "key", Value: "value"})
	assert.Error(t, err)
}
