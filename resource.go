package gotel

import (
	"context"
	"errors"
	"log"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"golang.org/x/mod/semver"
)

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	if !semver.IsValid(serviceVersion) {
		return nil, errors.New("serviceVersion is not a valid semver." +
			"The semver should look like vMAJOR[.MINOR[.PATCH[-PRERELEASE][+BUILD]]]")
	}
	res, err := resource.New(
		context.Background(),
		// Discover and provide attributes from OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables.
		resource.WithFromEnv(),
		// Discover and provide information about the OpenTelemetry SDK used.
		resource.WithTelemetrySDK(),
		// Discover and provide process information.
		resource.WithProcess(),
		// Discover and provide OS information.
		resource.WithOS(),
		// Discover and provide container information.
		resource.WithContainer(),
		// Discover and provide host information.
		resource.WithHost(),
	)
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		log.Println(err) // Log non-fatal issues.
	} else if err != nil {
		return nil, err
	}
	return resource.Merge(res,
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}
