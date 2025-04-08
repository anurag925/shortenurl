CREATE TABLE short_urls (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  short_code TEXT NOT NULL UNIQUE,
  long_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP NULL,
  visit_count INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_short_urls_short_code ON short_urls (short_code);