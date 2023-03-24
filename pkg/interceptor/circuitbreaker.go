package interceptor

import (
	"context"
	"errors"

	"github.com/mercari/go-circuitbreaker"
	"github.com/mgjules/minion/pkg/logger"
	"google.golang.org/grpc"
)

// CircuitBreaker is a client interceptor which wraps around circuitbreaker.CircuitBreaker.
type CircuitBreaker struct {
	cb     *circuitbreaker.CircuitBreaker
	logger *logger.Logger
}

// NewCircuitBreaker returns a new circuit breaker interceptor.
func NewCircuitBreaker(cb *circuitbreaker.CircuitBreaker, logger *logger.Logger) *CircuitBreaker {
	return &CircuitBreaker{cb, logger}
}

// UnaryClient returns a client interceptor function to implement circuit breaking on Unary RPC requests.
func (c *CircuitBreaker) UnaryClient() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		_, err := c.cb.Do(ctx, func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}

			return nil, nil
		})

		if errors.Is(err, circuitbreaker.ErrOpen) {
			c.logger.Ctx(ctx).Debugw("circuit is open")
		}

		return err
	}
}
