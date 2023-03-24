package domain

import (
	"context"
)

// WordRepository abstracts away the storage adapter.
type WordRepository interface {
	AddWord(ctx context.Context, word string) (*Word, error)
	ListWords(ctx context.Context, params ListWordsParams) ([]Word, error)
	RandomWord(ctx context.Context) (*Word, error)

	Close() error
}

// WordSearchRepository abstracts away the search adapter.
type WordSearchRepository interface {
	IndexWord(ctx context.Context, word Word) error
	SearchWordsByPrefix(ctx context.Context, prefix string) ([]int, error)

	Close() error
}

// ListWordsParams holds a list of attributes that words can be filtered by.
type ListWordsParams struct {
	IDs []int
}

// Word holds the attributes of a word.
type Word struct {
	ID   int
	Word string
}
