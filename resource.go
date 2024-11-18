package initialize

import (
	"errors"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"golang.org/x/mod/semver"
)

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	if !semver.IsValid(serviceVersion) {
		return nil, errors.New("serviceVersion is not a valid semver. The semver should look like vMAJOR[.MINOR[.PATCH[-PRERELEASE][+BUILD]]]")
	}
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}
