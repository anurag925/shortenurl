package repositories_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/anurag/shortenurl/internal/db/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func setupTestDB(t *testing.T) *bun.DB {
	sqlite, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	assert.NoError(t, err)

	db := bun.NewDB(sqlite, sqlitedialect.New())
	_, err = db.NewCreateTable().Model((*models.ShortURL)(nil)).Exec(context.Background())
	assert.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, db.Close())
	})

	return db
}

func TestShortURLRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewShortURLRepository(db)
	ctx := context.Background()

	t.Run("Create and Find", func(t *testing.T) {
		shortURL := &models.ShortURL{
			ShortCode: "abc123",
			LongURL:   "https://example.com",
		}

		err := repo.Create(ctx, shortURL)
		assert.NoError(t, err)
		assert.NotZero(t, shortURL.ID)

		found, err := repo.FindByShortCode(ctx, "abc123")
		assert.NoError(t, err)
		assert.Equal(t, shortURL.ID, found.ID)
		assert.Equal(t, shortURL.LongURL, found.LongURL)
	})

	t.Run("IncrementVisitCount", func(t *testing.T) {
		shortURL := &models.ShortURL{
			ShortCode: "def456",
			LongURL:   "https://example.org",
		}
		assert.NoError(t, repo.Create(ctx, shortURL))

		assert.NoError(t, repo.IncrementVisitCount(ctx, shortURL.ID))

		updated, err := repo.FindByShortCode(ctx, "def456")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), updated.VisitCount)
	})

	t.Run("IncrementVisitCount 2", func(t *testing.T) {
		shortURL := &models.ShortURL{
			ShortCode: "def456",
			LongURL:   "https://example.org",
		}
		assert.NoError(t, repo.Create(ctx, shortURL))

		assert.NoError(t, repo.IncrementVisitCount(ctx, shortURL.ID))
		assert.NoError(t, repo.IncrementVisitCount(ctx, shortURL.ID))

		updated, err := repo.FindByShortCode(ctx, "def456")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), updated.VisitCount)
	})
}
