package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"time"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/anurag/shortenurl/internal/db/repositories"
)

type shortURLServiceImpl struct {
	repo repositories.ShortURLRepository
}

func NewShortURLService(repo repositories.ShortURLRepository) ShortURLService {
	return &shortURLServiceImpl{repo: repo}
}

func (s *shortURLServiceImpl) ShortenURL(ctx context.Context, longURL string, userID *int64, customAlias *string, expiresAt *time.Time) (*models.ShortURL, error) {
	shortCode := ""
	if customAlias != nil {
		shortCode = *customAlias
	} else {
		shortCode = generateRandomShortCode(6)
	}

	shortURL := &models.ShortURL{
		ShortCode: shortCode,
		LongURL:   longURL,
		ExpiresAt: expiresAt,
	}
	if userID != nil {
		shortURL.UserID = userID
	}

	if err := s.repo.Create(ctx, shortURL); err != nil {
		return nil, err
	}

	return shortURL, nil
}

func (s *shortURLServiceImpl) GetOriginalURL(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	shortURL, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	if shortURL.ExpiresAt != nil && shortURL.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("URL has expired")
	}
	slog.InfoContext(ctx, "shortURL is being called", "long_url", shortURL.LongURL)
	if err := s.repo.IncrementVisitCount(ctx, shortURL.ID); err != nil {
		return nil, err
	}

	return shortURL, nil
}

func generateRandomShortCode(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}
