package dto

import "time"

type ShortenRequest struct {
	LongURL     string     `json:"long_url" validate:"required,url"`
	CustomAlias *string    `json:"custom_alias,omitempty" validate:"omitempty,alphanum,min=3,max=20"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type ShortenResponse struct {
	ShortCode string     `json:"short_code"`
	LongURL   string     `json:"long_url"`
	ShortURL  string     `json:"short_url"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
