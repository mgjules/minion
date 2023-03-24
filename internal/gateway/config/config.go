package config

// Config holds the gateway service configuration.
type Config struct {
	Debug             bool   `envconfig:"DEBUG"`
	Host              string `envconfig:"HOST" default:"0.0.0.0"`
	Port              int    `envconfig:"PORT" default:"9000"`
	OTLPEndpoint      string `envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT" default:"0.0.0.0:4317"`
	WordsGRPCEndpoint string `envconfig:"WORDS_GRPC_ENDPOINT" default:"0.0.0.0:9001"`
}
