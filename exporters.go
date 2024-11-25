package gotel

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func setupExporters(ctx context.Context, destination Destination) (
	trace.SpanExporter, metric.Exporter, log.Exporter, error,
) {
	switch destination {
	case GRPC:
		return grpcExporters(ctx)
	case HTTP:
		return httpExporters(ctx)
	case STDOUT:
		return stdExporters()
	default:
		return nil, nil, nil, errors.New("how did you manage to get outside the exhaustive list?")
	}
}

func httpExporters(ctx context.Context) (trace.SpanExporter, metric.Exporter, log.Exporter, error) {
	traceExporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	metricExporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	logExporter, err := otlploghttp.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return traceExporter, metricExporter, logExporter, nil
}

func stdExporters() (trace.SpanExporter, metric.Exporter, log.Exporter, error) {
	traceExporter, err := stdouttrace.New()
	if err != nil {
		return nil, nil, nil, err
	}
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, nil, nil, err
	}
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, nil, nil, err
	}
	return traceExporter, metricExporter, logExporter, nil
}

func grpcExporters(ctx context.Context) (trace.SpanExporter, metric.Exporter, log.Exporter, error) {
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	metricExporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	logExporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return traceExporter, metricExporter, logExporter, nil
}
