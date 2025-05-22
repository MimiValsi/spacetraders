-- name: registerToken :one
INSERT INTO tokens (id)
VALUES ($1);
--
