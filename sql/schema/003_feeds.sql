-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  updated_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  last_fetched_at TIMESTAMP,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;