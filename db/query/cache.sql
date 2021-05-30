-- name: Set :one
-- postgres has implicit tansaction for single statement
INSERT INTO cache (
    key, 
    value, 
    ttl
)
VALUES (
    $1, $2, $3
)
ON CONFLICT (key) 
DO UPDATE SET value = $2, ttl = $3 RETURNING *;

-- name: Get :one
SELECT * from cache 
WHERE key = $1 LIMIT 1;