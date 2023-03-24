package adapter

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/mgjules/minion/internal/words/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ domain.WordSearchRepository = &WordBleveSearchRepository{}

// WordBleveSearchRepository is an in-memory implementation of the domain.WordSearchRepository interface.
type WordBleveSearchRepository struct {
	index  bleve.Index
	tracer trace.Tracer
}

// NewWordBleveSearchRepository creates a new WordMemorySearchRepository.
func NewWordBleveSearchRepository(path string, tracer trace.Tracer) (*WordBleveSearchRepository, error) {
	if tracer == nil {
		return nil, errors.New("tracer must not be nil")
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(path, mapping)
	if err != nil {
		if !errors.Is(err, bleve.ErrorIndexPathExists) {
			return nil, fmt.Errorf("new bleve index: %w", err)
		}

		index, err = bleve.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open bleve index: %w", err)
		}
	}

	return &WordBleveSearchRepository{
		index:  index,
		tracer: tracer,
	}, nil
}

// Close closes any resources associated with the repository.
func (r *WordBleveSearchRepository) Close() error {
	return r.index.Close()
}

// IndexWord indexes a new word in the bleve.Index.
func (r *WordBleveSearchRepository) IndexWord(ctx context.Context, word domain.Word) error {
	_, span := r.tracer.Start(ctx, "WordBleveSearchRepository.IndexWord")
	span.SetAttributes(attribute.Key("word.ID").Int(word.ID))
	defer span.End()

	return r.index.Index(strconv.Itoa(word.ID), word)
}

// SearchWordsByPrefix returns a list of word ids with a specific prefix in the bleve.Index.
func (r *WordBleveSearchRepository) SearchWordsByPrefix(ctx context.Context, prefix string) ([]int, error) {
	_, span := r.tracer.Start(ctx, "WordBleveSearchRepository.SearchWordsByPrefix")
	span.SetAttributes(attribute.Key("prefix").String(prefix))
	defer span.End()

	prefixQuery := bleve.NewPrefixQuery(prefix)
	prefixQuery.SetField("Word")
	searchRequest := bleve.NewSearchRequest(prefixQuery)
	searchRequest.SortBy([]string{"ID"})

	searchResult, err := r.index.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("search index: %w", err)
	}

	ids := make([]int, len(searchResult.Hits))
	for i, doc := range searchResult.Hits {
		id, err := strconv.Atoi(doc.ID)
		if err != nil {
			continue
		}

		ids[i] = id
	}

	return ids, nil
}
