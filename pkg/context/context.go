package context

import (
	"whai/internal/ai"
	"whai/internal/config"
	"whai/internal/shell"
	"whai/pkg/utils/logger"
	"whai/pkg/utils/ui"
)

type Context struct {
	Logger             logger.Logger
	UI                 ui.UI
	Config             config.Config
	ConfigCommander    config.Commander
	AIResponseProvider ai.ResponseProvider
	ShellProvider      shell.Provider
}
