package port_test

import (
	context "context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mgjules/minion/internal/protobuf/words"
	"github.com/mgjules/minion/internal/words/adapter"
	"github.com/mgjules/minion/internal/words/port"
	"github.com/mgjules/minion/pkg/cache"
	"github.com/mgjules/minion/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/global"
)

const (
	maxKeys     = 1e7               // Num keys to track frequency of (10M).
	maxCost     = 100 * 1000 * 1024 // Maximum cost of cache (100MB in bytes).
	testTimeout = 10 * time.Second
)

func TestWordsGrpcServiceServer(t *testing.T) {
	t.Parallel()

	tracer := otel.Tracer("test-tracer")

	repo, err := adapter.NewWordMemoryRepository(tracer)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	t.Cleanup(func() {
		err = repo.Close()
		assert.NoError(t, err)
	})

	id := uuid.New().String()
	searchRepo, err := adapter.NewWordBleveSearchRepository(id, tracer)
	assert.NoError(t, err)
	assert.NotNil(t, repo)
	t.Cleanup(func() {
		err = searchRepo.Close()
		assert.NoError(t, err)
		err = os.RemoveAll(id)
		assert.NoError(t, err)
	})

	cache, err := cache.New(maxKeys, maxCost)
	assert.NoError(t, err)
	assert.NotNil(t, cache)

	logger, err := logger.New(false)
	assert.NoError(t, err)
	assert.NotNil(t, logger)

	meter := global.Meter("meter-tracer")

	serviceServer, err := port.NewWordsGrpcServiceServer(
		repo, searchRepo, cache, meter, logger,
	)
	assert.NoError(t, err)
	assert.NotNil(t, serviceServer)

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("add word", func(t *testing.T) { //nolint:paralleltest
		resp, err := serviceServer.AddWord(ctx, &words.AddWordRequest{
			Word: "helloworld",
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), resp.GetId())
		assert.Equal(t, "helloworld", resp.GetWord())

		resp, err = serviceServer.AddWord(ctx, &words.AddWordRequest{
			Word: "helloworlds",
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(2), resp.GetId())
		assert.Equal(t, "helloworlds", resp.GetWord())

		resp, err = serviceServer.AddWord(ctx, &words.AddWordRequest{
			Word: "hola",
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(3), resp.GetId()) //nolint:revive
		assert.Equal(t, "hola", resp.GetWord())

		resp, err = serviceServer.AddWord(ctx, &words.AddWordRequest{
			Word: "holas",
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(4), resp.GetId()) //nolint:revive
		assert.Equal(t, "holas", resp.GetWord())
	})

	t.Run("search word", func(t *testing.T) { //nolint:paralleltest
		resp, err := serviceServer.SearchWord(ctx, &words.SearchWordRequest{
			Query: "hello",
		})
		assert.NoError(t, err)
		assert.Equal(t, []string{"helloworld", "helloworlds"}, resp.GetWords())

		resp, err = serviceServer.SearchWord(ctx, &words.SearchWordRequest{
			Query: "ho",
		})
		assert.NoError(t, err)
		assert.Equal(t, []string{"hola", "holas"}, resp.GetWords())

		resp, err = serviceServer.SearchWord(ctx, &words.SearchWordRequest{
			Query: "h",
		})
		assert.NoError(t, err)
		assert.Equal(t, []string{"helloworld", "helloworlds", "hola", "holas"}, resp.GetWords()) //nolint:revive
	})

	t.Run("randomize word", func(t *testing.T) { //nolint:paralleltest
		resp, err := serviceServer.RandomWord(ctx, &words.RandomWordRequest{})
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.GetWord())
	})
}
