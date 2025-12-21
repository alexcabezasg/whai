package logger

import "github.com/pterm/pterm"

type Logger interface {
	Info(str string)
	Error(str string)
}

type DefaultLogger struct{}

func (DefaultLogger) Info(str string) {
	pterm.DefaultLogger.Info(str)
}

func (DefaultLogger) Error(str string) {
	pterm.DefaultLogger.Error(str)
}

func NewLogger() Logger {
	return DefaultLogger{}
}
