package adapter_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mgjules/minion/internal/words/adapter"
	"github.com/mgjules/minion/internal/words/domain"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
)

func TestWordBleveSearchRepository(t *testing.T) {
	t.Parallel()

	tracer := otel.Tracer("test-tracer")

	id := uuid.New().String()
	repo, err := adapter.NewWordBleveSearchRepository(id, tracer)
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	t.Cleanup(func() {
		err = repo.Close()
		assert.NoError(t, err)
		err = os.RemoveAll(id)
		assert.NoError(t, err)
	})

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("index word", func(t *testing.T) { //nolint:paralleltest
		err := repo.IndexWord(ctx, domain.Word{ID: 1, Word: "helloworld"})
		assert.NoError(t, err)

		err = repo.IndexWord(ctx, domain.Word{ID: 2, Word: "helloworlds"})
		assert.NoError(t, err)

		err = repo.IndexWord(ctx, domain.Word{ID: 3, Word: "hola"}) //nolint:revive
		assert.NoError(t, err)

		err = repo.IndexWord(ctx, domain.Word{ID: 4, Word: "holas"}) //nolint:revive
		assert.NoError(t, err)
	})

	t.Run("search words by prefix", func(t *testing.T) { //nolint:paralleltest
		ids, err := repo.SearchWordsByPrefix(ctx, "hello")
		assert.NoError(t, err)
		assert.Len(t, ids, 2)

		ids, err = repo.SearchWordsByPrefix(ctx, "holas")
		assert.NoError(t, err)
		assert.Len(t, ids, 1)

		ids, err = repo.SearchWordsByPrefix(ctx, "h")
		assert.NoError(t, err)
		assert.Len(t, ids, 4) //nolint:revive
	})
}

func BenchmarkWordBleveSearchRepository(b *testing.B) {
	tracer := otel.Tracer("test-tracer")

	id := uuid.New().String()
	repo, err := adapter.NewWordBleveSearchRepository(id, tracer)
	assert.NoError(b, err)
	assert.NotNil(b, repo)

	b.Cleanup(func() {
		err = repo.Close()
		assert.NoError(b, err)
		err = os.RemoveAll(id)
		assert.NoError(b, err)
	})

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	b.Cleanup(func() {
		cancel()
	})

	b.ResetTimer()

	b.Run("index word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			repo.IndexWord(ctx, domain.Word{ID: 1, Word: "helloworld"})
		}
	})

	b.Run("search words by prefix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			repo.SearchWordsByPrefix(ctx, "hello")
		}
	})
}
