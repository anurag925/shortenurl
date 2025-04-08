package repositories

import (
	"context"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/uptrace/bun"
)

type shortURLRepositoryBun struct {
	db *bun.DB
}

func NewShortURLRepository(db *bun.DB) ShortURLRepository {
	return &shortURLRepositoryBun{db: db}
}

func (r *shortURLRepositoryBun) Create(ctx context.Context, shortURL *models.ShortURL) error {
	_, err := r.db.NewInsert().
		Model(shortURL).
		Exec(ctx)
	return err
}

func (r *shortURLRepositoryBun) FindByShortCode(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	var shortURL models.ShortURL
	err := r.db.NewSelect().
		Model(&shortURL).
		Where("short_code = ?", shortCode).
		Scan(ctx)
	return &shortURL, err
}

func (r *shortURLRepositoryBun) IncrementVisitCount(ctx context.Context, id int64) error {
	_, err := r.db.NewUpdate().
		Model((*models.ShortURL)(nil)).
		Set("visit_count = visit_count + 1").
		Where("id = ?", id).
		Exec(ctx)
	return err
}
