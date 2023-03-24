package port

import (
	context "context"
	"errors"
	"fmt"
	"time"

	"github.com/mgjules/minion/internal/protobuf/words"
	"github.com/mgjules/minion/internal/words/domain"
	"github.com/mgjules/minion/pkg/cache"
	"github.com/mgjules/minion/pkg/logger"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var searchWordCacheKey = "SearchWord_"

// WordsGrpcServiceServer is an implementation of WordsServiceServer interface.
type WordsGrpcServiceServer struct {
	repo        domain.WordRepository
	searchRepo  domain.WordSearchRepository
	cache       *cache.Cache
	logger      *logger.Logger
	wordCounter instrument.Float64Counter
}

// NewWordsGrpcServiceServer creates a new WordsGrpcServiceServer.
func NewWordsGrpcServiceServer(
	repo domain.WordRepository,
	searchRepo domain.WordSearchRepository,
	cache *cache.Cache,
	meter metric.Meter,
	logger *logger.Logger,
) (*WordsGrpcServiceServer, error) {
	if repo == nil {
		return nil, errors.New("repo must not be nil")
	}
	if searchRepo == nil {
		return nil, errors.New("search repo must not be nil")
	}
	if cache == nil {
		return nil, errors.New("cache must not be nil")
	}
	if meter == nil {
		return nil, errors.New("meter must not be nil")
	}
	if logger == nil {
		return nil, errors.New("logger must not be nil")
	}

	wordCounter, err := meter.Float64Counter(
		"word_count",
		instrument.WithDescription("The number of words"),
	)
	if err != nil {
		return nil, fmt.Errorf("new word count: %w", err)
	}

	return &WordsGrpcServiceServer{
		repo:        repo,
		searchRepo:  searchRepo,
		cache:       cache,
		logger:      logger,
		wordCounter: wordCounter,
	}, nil
}

// AddWord implements words.WordsServiceServer.
func (s *WordsGrpcServiceServer) AddWord(
	ctx context.Context,
	req *words.AddWordRequest,
) (*words.AddWordResponse, error) {
	reqWord := req.GetWord()
	word, err := s.repo.AddWord(ctx, reqWord)
	if err != nil {
		s.logger.Ctx(ctx).Errorw("cannot add word", "word", reqWord, "error", err.Error())

		return nil, status.Errorf(codes.Internal, "cannot add word '%s'", reqWord)
	}

	if err = s.searchRepo.IndexWord(ctx, *word); err != nil {
		s.logger.Ctx(ctx).Errorw("cannot index word", "word", reqWord, "error", err.Error())

		return nil, status.Errorf(codes.Internal, "cannot index word '%s'", reqWord)
	}

	s.wordCounter.Add(ctx, 1)

	s.cache.Clear()

	return &words.AddWordResponse{
		Id:   int64(word.ID),
		Word: word.Word,
	}, nil
}

// RandomWord implements words.WordsServiceServer.
func (s *WordsGrpcServiceServer) RandomWord(
	ctx context.Context,
	_ *words.RandomWordRequest,
) (*words.RandomWordResponse, error) {
	word, err := s.repo.RandomWord(ctx)
	if err != nil {
		s.logger.Ctx(ctx).Errorw("get random word", "error", err.Error())

		return nil, status.Error(codes.Internal, "cannot get random word")
	}

	if word == nil {
		return nil, status.Error(codes.NotFound, "no random word found")
	}

	return &words.RandomWordResponse{
		Id:   int64(word.ID),
		Word: word.Word,
	}, nil
}

// SearchWord implements words.WordsServiceServer.
func (s *WordsGrpcServiceServer) SearchWord(
	ctx context.Context,
	req *words.SearchWordRequest,
) (*words.SearchWordResponse, error) {
	query := req.GetQuery()
	val, ok := s.cache.Get(searchWordCacheKey + query)
	if ok {
		respWords, ok := val.([]string)
		if ok {
			s.logger.Ctx(ctx).Debugw("retrieving words from cache", "words", respWords)

			return &words.SearchWordResponse{Words: respWords}, nil
		}
	}

	ids, err := s.searchRepo.SearchWordsByPrefix(ctx, query)
	if err != nil {
		s.logger.Ctx(ctx).Errorw("cannot search word by query", "query", query, "error", err.Error())

		return nil, status.Errorf(codes.Internal, "cannot search word by query '%s'", query)
	}

	if len(ids) == 0 {
		return &words.SearchWordResponse{}, nil
	}

	rWords, err := s.repo.ListWords(ctx, domain.ListWordsParams{IDs: ids})
	if err != nil {
		s.logger.Ctx(ctx).Errorw("cannot list words from ids", "ids", ids, "error", err.Error())

		return nil, status.Error(codes.Internal, "cannot list words from ids")
	}

	respWords := make([]string, len(rWords))
	for i, w := range rWords {
		respWords[i] = w.Word
	}

	s.cache.SetWithTTL(searchWordCacheKey+req.Query, respWords, 0, 1*time.Minute)

	return &words.SearchWordResponse{Words: respWords}, nil
}
