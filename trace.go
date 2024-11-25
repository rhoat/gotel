package gotel

import (
	"go.opentelemetry.io/otel/sdk/trace"
)

func newTraceProvider(cfg Config) *trace.TracerProvider {
	traceProvider := trace.NewTracerProvider(
		cfg.TracerProviderOption...,
	)
	return traceProvider
}
