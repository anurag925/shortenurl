package service_test

import (
	"context"
	"testing"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/anurag/shortenurl/internal/db/repositories"
	"github.com/anurag/shortenurl/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
	repositories.ShortURLRepository
}

func (m *mockRepository) Create(ctx context.Context, shortURL *models.ShortURL) error {
	args := m.Called(ctx, shortURL)
	return args.Error(0)
}

func (m *mockRepository) FindByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	args := m.Called(ctx, shortCode)
	return args.Get(0).(*models.ShortURL), args.Error(1)
}

func (m *mockRepository) IncrementVisitCount(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestShortURLService(t *testing.T) {
	ctx := context.Background()
	repo := new(mockRepository)
	service := service.NewShortURLService(repo)

	t.Run("ShortenURL", func(t *testing.T) {
		longURL := "https://example.com"
		expectedShortURL := &models.ShortURL{
			ShortCode: "abc123",
			LongURL:   longURL,
		}

		repo.On("Create", ctx, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.ShortURL)
			*arg = *expectedShortURL
		})

		result, err := service.ShortenURL(ctx, longURL, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedShortURL.ShortCode, result.ShortCode)
		assert.Equal(t, expectedShortURL.LongURL, result.LongURL)
	})

	t.Run("GetOriginalURL", func(t *testing.T) {
		shortCode := "abc123"
		expectedShortURL := &models.ShortURL{
			ID:        1,
			ShortCode: shortCode,
			LongURL:   "https://example.com",
		}

		repo.On("FindByShortCode", ctx, shortCode).Return(expectedShortURL, nil)
		repo.On("IncrementVisitCount", ctx, expectedShortURL.ID).Return(nil)

		result, err := service.GetOriginalURL(ctx, shortCode)
		assert.NoError(t, err)
		assert.Equal(t, expectedShortURL.LongURL, result.LongURL)
	})
}
