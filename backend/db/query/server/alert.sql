
-- name: CreateAlert :one
INSERT INTO alerts (host, severity, content)
VALUES ($1, $2, $3)
RETURNING id, host, severity, content, time;

-- name: GetAlertsByHost :many
SELECT id, host, severity, content, time
FROM alerts 
WHERE host = $1
ORDER BY time DESC;

-- name: GetAllAlerts :many
SELECT id, host, severity, content, time
FROM alerts 
ORDER BY time DESC
LIMIT $1;

-- name: GetRecentAlerts :many
SELECT id, host, severity, content, time
FROM alerts 
WHERE time > $1
ORDER BY time DESC;

-- name: DeleteOldAlerts :exec
DELETE FROM alerts 
WHERE time < $1;
