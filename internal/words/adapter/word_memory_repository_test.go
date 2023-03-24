package adapter_test

import (
	"context"
	"testing"
	"time"

	"github.com/mgjules/minion/internal/words/adapter"
	"github.com/mgjules/minion/internal/words/domain"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
)

const testTimeout = 10 * time.Second

func TestWordMemoryRepository(t *testing.T) {
	t.Parallel()

	tracer := otel.Tracer("test-tracer")

	repo, err := adapter.NewWordMemoryRepository(tracer)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	t.Cleanup(func() {
		err := repo.Close()
		assert.NoError(t, err)
	})

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("add word", func(t *testing.T) { //nolint:paralleltest
		word, err := repo.AddWord(ctx, "helloworld")
		assert.NoError(t, err)
		assert.NotNil(t, word)
		assert.Equal(t, 1, word.ID)
		assert.Equal(t, "helloworld", word.Word)

		word, err = repo.AddWord(ctx, "helloworlds")
		assert.NoError(t, err)
		assert.NotNil(t, word)
		assert.Equal(t, 2, word.ID)
		assert.Equal(t, "helloworlds", word.Word)

		word, err = repo.AddWord(ctx, "hola")
		assert.NoError(t, err)
		assert.NotNil(t, word)
		assert.Equal(t, 3, word.ID) //nolint:revive
		assert.Equal(t, "hola", word.Word)

		word, err = repo.AddWord(ctx, "holas")
		assert.NoError(t, err)
		assert.NotNil(t, word)
		assert.Equal(t, 4, word.ID) //nolint:revive
		assert.Equal(t, "holas", word.Word)
	})

	t.Run("list words", func(t *testing.T) { //nolint:paralleltest
		words, err := repo.ListWords(ctx, domain.ListWordsParams{})
		assert.NoError(t, err)
		assert.Len(t, words, 4) //nolint:revive

		words, err = repo.ListWords(ctx, domain.ListWordsParams{IDs: []int{2, 3, 5, 6}}) //nolint:revive
		assert.NoError(t, err)
		assert.Len(t, words, 2)
	})
}

func BenchmarkWordMemoryRepository(b *testing.B) {
	tracer := otel.Tracer("test-tracer")

	repo, err := adapter.NewWordMemoryRepository(tracer)
	assert.NoError(b, err)
	assert.NotNil(b, repo)
	b.Cleanup(func() {
		err := repo.Close()
		assert.NoError(b, err)
	})

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	b.Cleanup(func() {
		cancel()
	})

	b.ResetTimer()

	b.Run("add word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			repo.AddWord(ctx, "helloworld")
		}
	})

	b.Run("list words", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			repo.ListWords(ctx, domain.ListWordsParams{})
		}
	})

	b.Run("list words by IDs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			repo.ListWords(ctx, domain.ListWordsParams{IDs: []int{1, 2, 3, 4}}) //nolint:revive
		}
	})
}
