package options

import (
	"errors"
)

func Run(args []string, opts []RunnableOption) error {
	var errs []error
	for _, arg := range args {
		IsSupported := false
		for _, opt := range opts {
			if opt.AcceptsInput(arg) {
				errs = append(errs, opt.Run(args))
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
