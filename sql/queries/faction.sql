-- name: registerFaction :one
INSERT INTO factions (symbol, name, description, headquarters, traits_id, is_recruiting)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
--

-- name: registerTraits :one
INSERT INTO traits (symbol, name, description)
VALUES ($1, $2, $3);
--
