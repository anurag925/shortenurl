package models

import (
	"time"

	"github.com/uptrace/bun"
)

type ShortURL struct {
	bun.BaseModel `bun:"table:short_urls"`

	ID         int64      `bun:"id,pk,autoincrement"`
	ShortCode  string     `bun:"short_code,unique,notnull"`
	LongURL    string     `bun:"long_url,notnull"`
	CreatedAt  time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	ExpiresAt  *time.Time `bun:"expires_at"`
	VisitCount int64      `bun:"visit_count,notnull,default:0"`
}