-- name: RegisterContract :one
INSERT INTO contracts (
  id,
  faction_symbol,
  type,
  accepted,
  fulfilled,
  deadline_to_accept,
  agent_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING cid;
--

-- name: RegisterTerms :one
INSERT INTO terms (deadline, contract_id)
VALUES ($1, $2)
RETURNING id;
--

-- name: RegisterPayment :exec
INSERT INTO payments (on_accepted, on_fulfilled, term_id)
VALUES ($1, $2, $3);
--

-- name: RegisterDeliver :exec
INSERT INTO delivers (
  trade_symbol,
  destination_symbol,
  units_required,
  units_fulfilled,
  term_id
)
VALUES ($1, $2, $3, $4, $5);
--
