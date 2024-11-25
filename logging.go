package gotel

import (
	"go.opentelemetry.io/otel/sdk/log"
)

func newLoggerProvider(cfg Config) *log.LoggerProvider {
	provider := log.NewLoggerProvider(
		cfg.LoggerProviderOption...,
	)
	return provider
}
