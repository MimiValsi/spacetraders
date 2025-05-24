-- name: RegisterAgent :one
INSERT INTO agents (
  account_id,
  token,
  symbol,
  headquarters,
  credits,
  starting_faction,
  ship_count
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
--

-- name: GetAgentToken :one
SELECT token FROM agents;
--

-- name: GetAgent :one
SELECT token, credits, headquarters FROM agents;
--
