// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const truncate = `-- name: Truncate :exec
TRUNCATE bar
`

func (q *Queries) Truncate(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncate)
	return err
}
