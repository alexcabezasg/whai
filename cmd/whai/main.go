package main

import (
	"os"
	"whai/internal/ai"
	"whai/internal/commands"
	"whai/internal/config"
	"whai/pkg/context"
	"whai/pkg/utils/logger"
	"whai/pkg/utils/ui"
)

func main() {
	args := os.Args[1:]
	configRetriever := config.NewRetriever()
	configCommander := config.NewCommander()

	var log logger.Logger

	err, configuration := configRetriever.Get()
	if err != nil && os.IsNotExist(err) {
		configuration = configuration.NewConfig()
		err := configCommander.Set(configuration)
		log = logger.NewLogger(configuration)
		log.Debug("Configuration file not found. Creating it from scratch.")
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	log = logger.NewLogger(configuration)

	ctx := context.Context{
		Config:             configuration,
		UI:                 ui.NewUI(),
		Logger:             log,
		ConfigCommander:    configCommander,
		AIResponseProvider: ai.NewAIResponseProvider(configuration),
	}

	err = commands.Run(args, ctx)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return
	}
}
