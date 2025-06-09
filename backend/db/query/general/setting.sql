-- name: CreateUser :one
INSERT INTO users (name, role, email, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING id, name, role, email;

-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE name = $1 AND password_hash = $2;

-- name: AddMacAccess :one
INSERT INTO mac_access_status (mac, status)
VALUES ($1, $2)
RETURNING id, mac, status, created_at, updated_at;

-- name: RemoveMacAccess :exec
DELETE FROM mac_access_status
WHERE mac = $1;
