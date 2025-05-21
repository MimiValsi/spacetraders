-- +goose Up
CREATE TABLE agents (
  account_id TEXT NOT NULL PRIMARY KEY,
  token TEXT NOT NULL,
  symbol TEXT NOT NULL,
  headquarters TEXT NOT NULL,
  credits BIGINT NOT NULL,
  starting_faction TEXT NOT NULL,
  ship_count INTEGER NOT NULL
);
--

-- +goose Down
DROP TABLE agents;
--
