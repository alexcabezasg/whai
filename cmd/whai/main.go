package main

import (
	"os"
	"whai/internal/commands"
	"whai/internal/config"
	"whai/pkg/context"
	"whai/pkg/utils/logger"
	"whai/pkg/utils/ui"
)

func main() {
	args := os.Args[1:]
	configRetriever := config.NewRetriever()

	err, configuration := configRetriever.Get()
	if err != nil {
		panic(err)
	}

	ctx := context.Context{
		Config:          configuration,
		UI:              ui.NewUI(),
		Logger:          logger.NewLogger(configuration),
		ConfigCommander: config.NewCommander(),
	}

	err = commands.Run(args, commands.GetAvailableCommands(), ctx)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return
	}
}
