package options

import (
	"errors"
	"whai/pkg/context"
)

func Run(args []string, opts []RunnableOption, ctx context.Context) error {
	var errs []error
	for _, arg := range args {
		IsSupported := false
		for _, opt := range opts {
			if opt.AcceptsInput(arg) {
				err := opt.Run(arg, ctx)
				if err != nil {
					errs = append(errs, err)
				} else {
					ctx.Logger.Debug("Option " + arg + " applied successfully.")
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
