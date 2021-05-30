// Code generated by sqlc. DO NOT EDIT.
// source: cache.sql

package db

import (
	"context"
	"database/sql"
)

const delete = `-- name: Delete :exec
DELETE from cache 
WHERE key = $1
`

func (q *Queries) Delete(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, delete, key)
	return err
}

const get = `-- name: Get :one
SELECT key, value, ttl from cache 
WHERE key = $1 LIMIT 1
`

func (q *Queries) Get(ctx context.Context, key string) (Cache, error) {
	row := q.db.QueryRowContext(ctx, get, key)
	var i Cache
	err := row.Scan(&i.Key, &i.Value, &i.Ttl)
	return i, err
}

const set = `-- name: Set :one
INSERT INTO cache (
    key, 
    value, 
    ttl
)
VALUES (
    $1, $2, $3
)
ON CONFLICT (key) 
DO UPDATE SET value = $2, ttl = $3 RETURNING key, value, ttl
`

type SetParams struct {
	Key   string        `json:"key"`
	Value string        `json:"value"`
	Ttl   sql.NullInt32 `json:"ttl"`
}

// postgres has implicit tansaction for single statement
func (q *Queries) Set(ctx context.Context, arg SetParams) (Cache, error) {
	row := q.db.QueryRowContext(ctx, set, arg.Key, arg.Value, arg.Ttl)
	var i Cache
	err := row.Scan(&i.Key, &i.Value, &i.Ttl)
	return i, err
}
