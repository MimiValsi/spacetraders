-- +goose Up
CREATE TABLE accounts (
  id TEXT NOT NULL,
  email TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);
--

-- +goose Down
DROP TABLE accounts;
--
