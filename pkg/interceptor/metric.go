package interceptor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"google.golang.org/grpc"
)

// Metric is a client interceptor which collects metric.
type Metric struct {
	meter metric.Meter
}

// NewMetric returns a new metric interceptor.
func NewMetric(meter metric.Meter) *Metric {
	return &Metric{meter}
}

// UnaryClient returns a client interceptor function to collect metrics on Unary RPC requests.
func (m *Metric) UnaryClient() (grpc.UnaryClientInterceptor, error) {
	requestCount, err := m.meter.Float64Counter(
		"request_count",
		instrument.WithDescription("The number of requests per method"),
	)
	if err != nil {
		return nil, fmt.Errorf("new request count: %w", err)
	}

	return func(
		ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		requestCount.Add(ctx, 1, attribute.Key("method").String(method))

		return invoker(ctx, method, req, reply, cc, opts...)
	}, nil
}
