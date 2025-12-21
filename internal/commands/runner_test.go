package commands

import (
	"errors"
	"testing"
	"whai/internal/config"
	"whai/pkg/utils/ui"

	"github.com/stretchr/testify/assert"
)

type DummyCommand Command

func (c DummyCommand) GetCommand() Command {
	return Command(c)
}

func (c DummyCommand) AcceptsInput(input string) bool {
	return c.Alias == input
}

func (c DummyCommand) Run(args []string, ui ui.UI, provider config.Provider) error {
	return errors.New("command executed")
}

func (c DummyCommand) New() DummyCommand {
	return DummyCommand{
		Alias:       "test",
		Description: "Its just a test",
	}
}

func TestRun(t *testing.T) {
	commands := []RunnableCommand{
		DummyCommand{}.New(),
	}

	err := Run([]string{"test"}, commands)
	assert.Error(t, err)
	assert.ErrorContainsf(t, err, "command executed", "")

	err = Run([]string{"other"}, commands)
	assert.ErrorContainsf(t, err, "Command other not found. Try 'whai help' for more information.", "")
}
