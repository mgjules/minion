package tracer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

// Setup setups the OpenTelemetry metric exporter & provider.
// Returns a cleanup function.
func Setup(ctx context.Context, prod bool, service, otlpEndpoint string) (func() error, error) {
	if otlpEndpoint == "" {
		return nil, errors.New("OTLP endpoint must not be empty")
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otlpEndpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	sctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	exporter, err := otlptrace.New(sctx, client)
	if err != nil {
		return nil, fmt.Errorf("new trace exporter: %w", err)
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

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	return func() error {
		newCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := provider.Shutdown(newCtx); err != nil {
			return fmt.Errorf("shutdown trace provider: %w", err)
		}

		return nil
	}, nil
}
