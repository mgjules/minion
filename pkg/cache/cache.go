package cache

import (
	"fmt"

	"github.com/DmitriyVTitov/size"
	"github.com/dgraph-io/ristretto"
)

const defaultBufferItems = 64

// Cache is a simple wrapper around ristretto.Cache.
type Cache struct {
	*ristretto.Cache
}

// New creates a new Cache.
func New(maxKeys, maxCost int64) (*Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: maxKeys,
		MaxCost:     maxCost,
		BufferItems: defaultBufferItems,
		Cost:        Coster,
	})
	if err != nil {
		return nil, fmt.Errorf("new ristretto cache: %w", err)
	}

	return &Cache{cache}, nil
}

// Coster returns the cost of a given value in bytes.
func Coster(v interface{}) int64 {
	s := size.Of(v)
	if s < 0 {
		return 1
	}

	return int64(s)
}
