-- name: RegisterAccount :exec
INSERT INTO accounts (id, email, created_at)
VALUES ($1, $2, $3);
--
