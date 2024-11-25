package gotel

import (
	"time"

	"github.com/rhoat/go-exercise/pkg/system"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	TracerProviderOption []trace.TracerProviderOption
	MetricProviderOption []metric.Option
	LoggerProviderOption []log.LoggerProviderOption
}

type CfgOptionFunc func(c *Config) error

var (
	traceBatchTimeout = 5 * time.Second
	metricInterval    = time.Minute
)

func NewConfig(
	traceExporter trace.SpanExporter,
	metricExporter metric.Exporter,
	logExporter log.Exporter,
	opts ...CfgOptionFunc,
) (*Config, error) {
	res, err := newResource(system.ApplicationName, system.BuildVersion)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		TracerProviderOption: []trace.TracerProviderOption{
			trace.WithBatcher(traceExporter,
				trace.WithBatchTimeout(traceBatchTimeout)),
			trace.WithResource(res),
		},
		MetricProviderOption: []metric.Option{
			metric.WithResource(res),
			metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(metricInterval))),
		},
		LoggerProviderOption: []log.LoggerProviderOption{log.WithResource(res),
			log.WithProcessor(log.NewBatchProcessor(logExporter))},
	}
	for _, opt := range opts {
		if err = opt(&cfg); err != nil {
			return nil, err
		}
	}
	return &cfg, nil
}
