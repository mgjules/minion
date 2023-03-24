package adapter

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/mgjules/minion/internal/words/domain"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ domain.WordRepository = &WordMemoryRepository{}

// WordMemoryRepository is an in-memory implementation of the domain.WordRepository interface.
type WordMemoryRepository struct {
	mu     sync.RWMutex
	words  map[int]domain.Word
	tracer trace.Tracer
}

// NewWordMemoryRepository creates a new WordMemoryRepository.
func NewWordMemoryRepository(tracer trace.Tracer) (*WordMemoryRepository, error) {
	if tracer == nil {
		return nil, errors.New("tracer must not be nil")
	}

	return &WordMemoryRepository{
		words:  make(map[int]domain.Word),
		tracer: tracer,
	}, nil
}

// Close closes any resources associated with the repository.
func (*WordMemoryRepository) Close() error {
	return nil
}

// AddWord adds a new word in the words map.
// TODO(mike): Cater for duplicates?
func (r *WordMemoryRepository) AddWord(ctx context.Context, word string) (*domain.Word, error) {
	_, span := r.tracer.Start(ctx, "WordMemoryRepository.AddWord")
	span.SetAttributes(attribute.Key("word.Word").String(word))
	defer span.End()

	r.mu.Lock()
	defer r.mu.Unlock()

	// Sort ids.
	ids := lo.Keys(r.words)
	sort.Ints(ids)

	var lastID int
	if len(ids) != 0 {
		lastID = ids[len(ids)-1]
	}

	w := domain.Word{
		ID:   lastID + 1,
		Word: word,
	}

	r.words[w.ID] = w

	return &w, nil
}

// ListWords list all the words in the map.
func (r *WordMemoryRepository) ListWords(ctx context.Context, params domain.ListWordsParams) ([]domain.Word, error) {
	_, span := r.tracer.Start(ctx, "WordMemoryRepository.ListWords")
	if len(params.IDs) != 0 {
		span.SetAttributes(attribute.Key("params.IDs").IntSlice(params.IDs))
	}
	defer span.End()

	r.mu.RLock()
	defer r.mu.RUnlock()

	var words []domain.Word

	if len(params.IDs) != 0 {
		for _, id := range params.IDs {
			word, ok := r.words[id]
			if !ok {
				continue
			}
			words = append(words, word)
		}
	} else {
		words = lo.Values(r.words)
	}

	return words, nil
}

// RandomWord implements domain.WordRepository
func (r *WordMemoryRepository) RandomWord(ctx context.Context) (*domain.Word, error) {
	_, span := r.tracer.Start(ctx, "WordMemoryRepository.RandomWord")
	defer span.End()

	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.words) == 0 {
		return nil, errors.New("no words to randomize")
	}

	ids := lo.Keys(r.words)
	randomized := lo.Shuffle(ids)

	word, ok := r.words[randomized[0]]
	if !ok {
		// technically we should never get here.
		return nil, fmt.Errorf("word id '%d' not found", randomized)
	}

	return &word, nil
}
