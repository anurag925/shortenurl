package repositories

import (
	"context"

	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}

type userRepositoryBun struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepositoryBun{db: db}
}

func (r *userRepositoryBun) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *userRepositoryBun) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user := new(models.User)
	err := r.db.NewSelect().
		Model(user).
		Where("username = ?", username).
		Scan(ctx)
	return user, err
}
