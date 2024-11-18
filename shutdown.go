package initialize

import (
	"context"
	"errors"
)

type shutdownFunc func(context.Context) error

// shutdown calls cleanup functions registered via shutdownFuncs.
// The errors from the calls are joined.
// Each registered cleanup will be invoked once.
func generateShutdownFunc(shutdownFuncs []func(context.Context) error) shutdownFunc {
	return func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
}

func handleErrors(inErr error, shutdownFn shutdownFunc, ctx context.Context) error {
	return errors.Join(inErr, shutdownFn(ctx))
}
