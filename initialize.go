package gotel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
)

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func SetupOTelSDK(ctx context.Context, serviceName, serviceVersion string, options ...CfgOptionFunc) (shutdownFunc, error) {
	var shutdownFuncs []func(context.Context) error
	cfg := baseCfg
	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	for _, opt := range options {
		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}
	res, err := newResource(serviceName, serviceVersion)
	if err != nil {
		return nil, err
	}

	// Set up trace provider.
	tracerProvider, err := newTraceProvider(ctx, res, cfg)
	if err != nil {
		return nil, handleErrors(err, generateShutdownFunc(shutdownFuncs), ctx)
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider, err := newMeterProvider(ctx, res, cfg)
	if err != nil {
		return nil, handleErrors(err, generateShutdownFunc(shutdownFuncs), ctx)
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)
	// Set up log provider.
	loggerProvider, err := newLoggerProvider(ctx, res, cfg)
	if err != nil {
		panic(err)
	}

	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)
	return generateShutdownFunc(shutdownFuncs), nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
