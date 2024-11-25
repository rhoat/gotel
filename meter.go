package gotel

import (
	"go.opentelemetry.io/otel/sdk/metric"
)

func newMeterProvider(cfg Config) *metric.MeterProvider {
	meterProvider := metric.NewMeterProvider(
		cfg.MetricProviderOption...,
	)

	return meterProvider
}
