package logger

import (
	"whai/internal/config"

	"github.com/pterm/pterm"
)

type Logger interface {
	Info(str string)
	Error(str string)
	Debug(str string)
}

type DefaultLogger struct {
	Logger *pterm.Logger
}

func (l DefaultLogger) Info(str string) {
	l.Logger.Info(str)
}

func (l DefaultLogger) Error(str string) {
	l.Logger.Error(str)
}

func (l DefaultLogger) Debug(str string) {
	l.Logger.Debug(str)
}

func NewLogger(config config.Config) Logger {
	logLevel := pterm.LogLevelInfo

	if config.DebugMode {
		logLevel = pterm.LogLevelDebug
	}
	return DefaultLogger{
		Logger: pterm.DefaultLogger.WithLevel(logLevel),
	}
}
