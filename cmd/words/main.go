package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mgjules/minion/internal/protobuf/words"
	"github.com/mgjules/minion/internal/words/adapter"
	wconfig "github.com/mgjules/minion/internal/words/config"
	"github.com/mgjules/minion/internal/words/port"
	"github.com/mgjules/minion/pkg/cache"
	"github.com/mgjules/minion/pkg/config"
	"github.com/mgjules/minion/pkg/interceptor"
	"github.com/mgjules/minion/pkg/logger"
	"github.com/mgjules/minion/pkg/metric"
	"github.com/mgjules/minion/pkg/tracer"
	"github.com/mgjules/minion/pkg/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/global"
)

const (
	serviceName     = "words"
	maxKeys         = 1e7               // Num keys to track frequency of (10M).
	maxCost         = 100 * 1000 * 1024 // Maximum cost of cache (100MB in bytes).
	shutdownTimeout = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start service: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load[wconfig.Config](serviceName)
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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

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

	tracerCleanup, err := tracer.Setup(ctx, !cfg.Debug, serviceName, cfg.OTLPEndpoint)
	if err != nil {
		return fmt.Errorf("setup tracer: %w", err)
	}
	defer func() {
		if err = tracerCleanup(); err != nil {
			logger.Errorf("tracer cleanup: %v", err)
		}
	}()
	tracer := otel.Tracer(serviceName + "-tracer")

	wordMemRepo, err := adapter.NewWordMemoryRepository(tracer)
	if err != nil {
		return fmt.Errorf("new word in-memory repo: %w", err)
	}
	defer func() {
		if err = wordMemRepo.Close(); err != nil {
			logger.Errorf("word in-memory repo close: %v", err)
		}
	}()

	wordBleveSearchRepo, err := adapter.NewWordBleveSearchRepository("words.bleve", tracer)
	if err != nil {
		return fmt.Errorf("new word bleve search repo: %w", err)
	}
	defer func() {
		if err = wordBleveSearchRepo.Close(); err != nil {
			logger.Errorf("word bleve search repo close: %v", err)
		}
		// We remove the index path since we want to start fresh each time.
		if err = os.RemoveAll("words.bleve"); err != nil {
			logger.Errorf("word bleve search repo remove index path: %v", err)
		}
	}()

	cache, err := cache.New(maxKeys, maxCost)
	if err != nil {
		return fmt.Errorf("new cache: %w", err)
	}
	defer cache.Close()

	wordsGrpcServiceServer, err := port.NewWordsGrpcServiceServer(
		wordMemRepo,
		wordBleveSearchRepo,
		cache,
		meter,
		logger,
	)
	if err != nil {
		return fmt.Errorf("new word grpc service server: %w", err)
	}

	grpcServer, err := transport.NewGRPCServer(
		cfg.Host,
		cfg.Port,
		logger,
		interceptor.NewValidator(logger).Unary(),
	)
	if err != nil {
		return fmt.Errorf("new grpc server: %w", err)
	}

	words.RegisterWordsServiceServer(grpcServer, wordsGrpcServiceServer)

	errCh := make(chan error)
	go func() {
		if err = grpcServer.Start(); err != nil {
			errCh <- fmt.Errorf("start grpc server: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		// Proceed with graceful shutdown of grpc Server.
	}

	ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	return grpcServer.Stop(ctx)
}
