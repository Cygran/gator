// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: reset_users.sql

package database

import (
	"context"
)

const reset = `-- name: Reset :exec

DELETE FROM users
`

func (q *Queries) Reset(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, reset)
	return err
}
