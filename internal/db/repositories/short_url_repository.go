package repositories

import (
	"context"

	"github.com/anurag/shortenurl/internal/db/models"
)

type ShortURLRepository interface {
	Create(ctx context.Context, shortURL *models.ShortURL) error
	FindByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error)
	IncrementVisitCount(ctx context.Context, id int64) error
}
