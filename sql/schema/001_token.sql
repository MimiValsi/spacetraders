-- +goose Up
CREATE TABLE tokens(
  id TEXT NOT NULL
);
--

-- +goose Down
DROP TABLE tokens;
--
