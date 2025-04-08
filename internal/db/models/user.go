package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `bun:"id,pk,autoincrement"`
	Username     string    `bun:"username,unique,notnull"`
	PasswordHash string    `bun:"password_hash,notnull"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}
