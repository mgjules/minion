package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mercari/go-circuitbreaker"
	gconfig "github.com/mgjules/minion/internal/gateway/config"
	"github.com/mgjules/minion/internal/openapi"
	"github.com/mgjules/minion/internal/protobuf/words"
	"github.com/mgjules/minion/pkg/config"
	"github.com/mgjules/minion/pkg/interceptor"
	"github.com/mgjules/minion/pkg/logger"
	"github.com/mgjules/minion/pkg/metric"
	"github.com/mgjules/minion/pkg/tracer"
	"github.com/mgjules/minion/pkg/transport"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric/global"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	serviceName     = "gateway"
	shutdownTimeout = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start service: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load[gconfig.Config](serviceName)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger, err := logger.New(!cfg.Debug)
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}
	defer func() {
		if err = logger.Sync(); err != nil {
			logger.Errorf("logger sync: %v", err)
		}
	}()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	//nolint:revive
	cb := circuitbreaker.New(
		circuitbreaker.WithCounterResetInterval(time.Minute),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncThreshold(3)),
		circuitbreaker.WithOpenTimeout(2500*time.Millisecond),
		circuitbreaker.WithHalfOpenMaxSuccesses(3),
	)

	tracerCleanup, err := tracer.Setup(ctx, !cfg.Debug, serviceName, cfg.OTLPEndpoint)
	if err != nil {
		return fmt.Errorf("setup tracer: %w", err)
	}
	defer func() {
		if err = tracerCleanup(); err != nil {
			logger.Errorf("tracer cleanup: %v", err)
		}
	}()

	metricCleanup, err := metric.Setup(ctx, !cfg.Debug, serviceName, cfg.OTLPEndpoint)
	if err != nil {
		return fmt.Errorf("setup metric: %w", err)
	}
	defer func() {
		if err = metricCleanup(); err != nil {
			logger.Errorf("metric cleanup: %v", err)
		}
	}()
	meter := global.Meter(serviceName + "-meter")

	metricInterceptorUnaryClient, err := interceptor.NewMetric(meter).UnaryClient()
	if err != nil {
		return fmt.Errorf("new metric interceptor unary client: %w", err)
	}

	grpcDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				interceptor.NewCircuitBreaker(cb, logger).UnaryClient(),
				otelgrpc.UnaryClientInterceptor(),
				metricInterceptorUnaryClient,
			),
		),
	}

	if err = words.RegisterWordsServiceHandlerFromEndpoint(
		ctx,
		mux,
		cfg.WordsGRPCEndpoint,
		grpcDialOptions,
	); err != nil {
		return fmt.Errorf("register words service handler: %w", err)
	}

	httpServer, err := transport.NewHTTPServer(
		!cfg.Debug,
		serviceName,
		cfg.Host,
		cfg.Port,
		logger,
	)
	if err != nil {
		return fmt.Errorf("new http server: %w", err)
	}

	httpServer.StaticFS("/openapi", http.FS(openapi.Definitions))
	httpServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/openapi/words/words.swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
	httpServer.Group("/v1/*{grpc_gateway}").Any("", gin.WrapH(mux))

	errCh := make(chan error)
	go func() {
		if err = httpServer.Start(); err != nil {
			errCh <- fmt.Errorf("start http server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		// Proceed with graceful shutdown of http Server.
	}

	ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	return httpServer.Stop(ctx)
}
