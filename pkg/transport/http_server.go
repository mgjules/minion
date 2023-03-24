package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mgjules/minion/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	readTimeout       = 2 * time.Second
	writeTimeout      = 2 * time.Second
	idleTimeout       = 30 * time.Second
	readHeaderTimeout = 2 * time.Second
)

// HTTPServer is the main HTTP server.
type HTTPServer struct {
	http   *http.Server
	logger *logger.Logger
	addr   string

	*gin.Engine
}

// NewHTTPServer creates a new Server.
func NewHTTPServer(
	prod bool,
	service string,
	host string,
	port int,
	logger *logger.Logger,
) (*HTTPServer, error) {
	if logger == nil {
		return nil, errors.New("logger must not be nil")
	}

	if prod {
		gin.SetMode(gin.ReleaseMode)
	}

	w := logger.Writer()
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w

	s := HTTPServer{
		Engine: gin.Default(),
		addr:   fmt.Sprintf("%s:%d", host, port),
		logger: logger,
	}

	desugared := logger.Desugar()
	s.Use(ginzap.Ginzap(desugared.Logger, time.RFC3339, true))
	s.Use(ginzap.RecoveryWithZap(desugared.Logger, true))
	s.Use(otelgin.Middleware(service))

	s.http = &http.Server{
		Addr:              s.addr,
		Handler:           s.Engine,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &s, nil
}

// Start starts the server.
// It blocks until the server stops.
func (s *HTTPServer) Start() error {
	s.logger.Infof("Listening HTTP server on http://%s...", s.addr)

	if err := s.http.ListenAndServe(); err != nil {
		return fmt.Errorf("start: %w", err)
	}

	return nil
}

// Stop stops the server.
func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping HTTP server ...")

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("stop: %w", err)
	}

	return nil
}
