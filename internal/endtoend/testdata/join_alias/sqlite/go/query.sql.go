// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const aliasExpand = `-- name: AliasExpand :many
SELECT f.id, b.id, title
FROM foo f
JOIN bar b ON b.id = f.id
WHERE f.id = ?
`

type AliasExpandRow struct {
	ID    int64
	ID_2  int64
	Title sql.NullString
}

func (q *Queries) AliasExpand(ctx context.Context, id int64) ([]AliasExpandRow, error) {
	rows, err := q.db.QueryContext(ctx, aliasExpand, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AliasExpandRow
	for rows.Next() {
		var i AliasExpandRow
		if err := rows.Scan(&i.ID, &i.ID_2, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const aliasJoin = `-- name: AliasJoin :many
SELECT f.id, b.title
FROM foo f
JOIN bar b ON b.id = f.id
WHERE f.id = ?
`

type AliasJoinRow struct {
	ID    int64
	Title sql.NullString
}

func (q *Queries) AliasJoin(ctx context.Context, id int64) ([]AliasJoinRow, error) {
	rows, err := q.db.QueryContext(ctx, aliasJoin, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AliasJoinRow
	for rows.Next() {
		var i AliasJoinRow
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
