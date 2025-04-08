package service

import (
	"context"
	"time"

	"github.com/anurag/shortenurl/internal/db/models"
)

type ShortURLService interface {
	ShortenURL(ctx context.Context, longURL string, customAlias *string, expiresAt *time.Time) (*models.ShortURL, error)
	GetOriginalURL(ctx context.Context, shortCode string) (*models.ShortURL, error)
}
