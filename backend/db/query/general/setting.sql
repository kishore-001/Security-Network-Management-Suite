-- name: CreateUser :one
INSERT INTO users (name, role, email, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING id, name, role, email;

