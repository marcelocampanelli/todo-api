-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUsersByID :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (name, email, password, created_at, updated_at)
values ($1, $2, $3, $4, $5);

-- name: UpdateUser :exec
UPDATE users
SET name = $1, email = $2, password = $3, updated_at = $4
WHERE id = $5;