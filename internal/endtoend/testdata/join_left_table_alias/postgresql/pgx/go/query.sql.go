// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const joinBar = `-- name: JoinBar :one
SELECT f.id, info
FROM foo f
LEFT JOIN bar b ON b.foo_id = f.id
`

type JoinBarRow struct {
	ID   int64
	Info pgtype.Text
}

func (q *Queries) JoinBar(ctx context.Context) (JoinBarRow, error) {
	row := q.db.QueryRow(ctx, joinBar)
	var i JoinBarRow
	err := row.Scan(&i.ID, &i.Info)
	return i, err
}

const joinBarAlias = `-- name: JoinBarAlias :one
SELECT f.id, b.info
FROM foo f
LEFT JOIN bar b ON b.foo_id = f.id
`

type JoinBarAliasRow struct {
	ID   int64
	Info pgtype.Text
}

func (q *Queries) JoinBarAlias(ctx context.Context) (JoinBarAliasRow, error) {
	row := q.db.QueryRow(ctx, joinBarAlias)
	var i JoinBarAliasRow
	err := row.Scan(&i.ID, &i.Info)
	return i, err
}
