-- name: GetUserByName :one
SELECT password_hash, role FROM users WHERE name = $1;

