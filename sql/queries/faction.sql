-- name: RegisterFaction :one
INSERT INTO factions (symbol, name, description, headquarters, is_recruiting, agent_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
--

-- name: RegisterTraits :exec
INSERT INTO traits (symbol, name, description, faction_id)
VALUES ($1, $2, $3, $4);
--
