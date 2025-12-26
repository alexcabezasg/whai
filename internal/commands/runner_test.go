package commands

import (
	"testing"
	"whai/pkg/context"

	"github.com/stretchr/testify/assert"
)

type DummyCommand Command

func (c DummyCommand) GetCommand() Command {
	return Command(c)
}

func (c DummyCommand) AcceptsInput(input string) bool {
	return c.Alias == input
}

func (c DummyCommand) Run(args []string, ctx context.Context) error {
	return nil
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

	ctx := context.Context{}

	err := Run([]string{"test"}, ctx, commands...)
	assert.Nil(t, err)

	err = Run([]string{"other"}, ctx, commands...)
	assert.ErrorContainsf(t, err, "Command other not found. Try 'whai help' for more information.", "")
}
