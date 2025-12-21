package commands

type Command struct {
	Alias       string
	Description string
	SubCommands []Command
}

type RunnableCommand interface {
	Run(args []string) error
	AcceptsInput(input string) bool
}

func (c Command) AcceptsInput(input string) bool {
	return input == c.Alias
}
