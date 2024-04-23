// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const getAuthor = `-- name: GetAuthor :one
SELECT a.name
FROM authors a JOIN authors b ON a.id = b.id
ORDER BY name
`

func (q *Queries) GetAuthor(ctx context.Context) (string, error) {
	row := q.db.QueryRow(ctx, getAuthor)
	var name string
	err := row.Scan(&name)
	return name, err
}
