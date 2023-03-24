package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Load parses and returns a new Config.
func Load[T any](service string) (*T, error) {
	var cfg T
	if err := envconfig.Process(service, &cfg); err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	return &cfg, nil
}
