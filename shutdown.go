package gotel

import (
	"context"
	"errors"
)

var (
	ShutDown ShutdownFunc
)

type ShutdownFunc func(context.Context) error

// shutdown calls cleanup functions registered via shutdownFuncs.
// The errors from the calls are joined.
// Each registered cleanup will be invoked once.
func generateShutdownFunc(shutdownFuncs []func(context.Context) error) ShutdownFunc {
	return func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		return err
	}
}
