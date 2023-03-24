package metric

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// Setup setups the OpenTelemetry metric exporter & provider.
// Returns a cleanup function.
func Setup(ctx context.Context, prod bool, service string, otlpEndpoint string) (func() error, error) {
	if otlpEndpoint == "" {
		return nil, errors.New("OTLP endpoint must not be empty")
	}

	env := "development"
	if prod {
		env = "production"
	}

	res, err := resource.New(ctx,
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", env),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("new resource: %w", err)
	}

	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(otlpEndpoint),
	)
	if err != nil {
		return nil, fmt.Errorf("new metric exporter: %w", err)
	}

	provider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(
			metric.NewPeriodicReader(
				exporter,
				metric.WithInterval(2*time.Second),
			),
		),
	)
	global.SetMeterProvider(provider)

	return func() error {
		newCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := provider.Shutdown(newCtx); err != nil {
			return fmt.Errorf("shutdown metric provider: %w", err)
		}

		return nil
	}, nil
}
