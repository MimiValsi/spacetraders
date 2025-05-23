-- name: RegisterToken :exec
INSERT INTO tokens (id)
VALUES ($1);
--
