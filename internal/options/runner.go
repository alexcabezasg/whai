package options

import (
	"errors"
	"whai/internal/config"
	"whai/pkg/utils/logger"
)

func Run(args []string, opts []RunnableOption, provider config.Provider, log logger.Logger) error {
	var errs []error
	for _, arg := range args {
		IsSupported := false
		for _, opt := range opts {
			if opt.AcceptsInput(arg) {
				err := opt.Run(arg, provider)
				if err != nil {
					errs = append(errs, err)
				} else {
					log.Info("Option " + arg + " applied successfully.")
					IsSupported = true
					continue
				}
			}
		}
		if !IsSupported {
			errs = append(errs, errors.New("Argument "+arg+" not supported."))
		}
	}

	return errors.Join(errs...)
}
