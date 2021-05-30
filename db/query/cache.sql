-- name: Set :exec
-- postgres has implicit tansaction for single statement
INSERT INTO cache (
    cache_key, 
    value, 
    ttl
)
VALUES (
    ?, ?, ?
);

-- name: Get :one
SELECT * from cache 
WHERE cache_key = ? LIMIT 1;

-- name: Delete :exec
DELETE from cache 
WHERE cache_key = ?;