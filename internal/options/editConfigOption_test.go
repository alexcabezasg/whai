package options

import (
	"testing"
	"whai/internal/config"
	"whai/pkg/context"

	"github.com/stretchr/testify/assert"
)

type DummyConfigProvider struct{}

func (DummyConfigProvider) Set(config.Config) error {
	return nil
}

func TestEditConfigOption(t *testing.T) {
	ctx := context.Context{
		ConfigCommander: DummyConfigProvider{},
		Config:          config.Config{OnlySuggest: true},
	}

	err := EditConfigOption{}.Run("--only-suggest=false", ctx)
	assert.Nil(t, err)
}

func TestEditConfigOption_AcceptsInput(t *testing.T) {
	ctx := context.Context{
		ConfigCommander: DummyConfigProvider{},
		Config:          config.Config{OnlySuggest: true},
	}

	err := EditConfigOption{}.Run("-only-suggest=false", ctx)
	assert.Error(t, err)
}
