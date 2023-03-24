package transport

import (
	"context"
	"errors"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/mgjules/minion/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer is a simple wrapper around grpc.Server.
type GRPCServer struct {
	addr   string
	logger *logger.Logger

	*grpc.Server
}

// NewGRPCServer creates a new grpc server.
func NewGRPCServer(
	host string,
	port int,
	logger *logger.Logger,
	customUnaryInterceptors ...grpc.UnaryServerInterceptor,
) (*GRPCServer, error) {
	if logger == nil {
		return nil, errors.New("logger must not be nil")
	}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(logger.Desugar().Logger),
		grpc_recovery.UnaryServerInterceptor(),
		otelgrpc.UnaryServerInterceptor(),
	}
	unaryInterceptors = append(unaryInterceptors, customUnaryInterceptors...)

	s := &GRPCServer{
		addr:   fmt.Sprintf("%s:%d", host, port),
		logger: logger,
		Server: grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(
					unaryInterceptors...,
				),
			),
		),
	}

	reflection.Register(s)

	return s, nil
}

// Start starts the grpc server.
func (s *GRPCServer) Start() error {
	s.logger.Infof("Listening gRPC server on tcp://%s...", s.addr)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	if err := s.Server.Serve(lis); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

// Stop stops the grpc server.
func (s *GRPCServer) Stop(context.Context) error {
	s.logger.Info("Stopping gRPC server ...")

	s.Server.GracefulStop()

	return nil
}
