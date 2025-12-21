package options

import (
	"errors"
	"whai/internal/config"
)

func Run(args []string, opts []RunnableOption, provider config.Provider) error {
	var errs []error
	for _, arg := range args {
		IsSupported := false
		for _, opt := range opts {
			if opt.AcceptsInput(arg) {
				errs = append(errs, opt.Run(args, provider))
				IsSupported = true
				continue
			}
		}
		if !IsSupported {
			errs = append(errs, errors.New("Argument "+arg+" not supported."))
		}
	}

	return errors.Join(errs...)
}
