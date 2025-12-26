package context

import (
	"whai/internal/config"
	"whai/pkg/utils/logger"
	"whai/pkg/utils/ui"
)

type Context struct {
	Logger          logger.Logger
	UI              ui.UI
	Config          config.Config
	ConfigCommander config.Commander
}
