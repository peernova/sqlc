// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const findByID = `-- name: FindByID :many
SELECT id, name FROM users WHERE $1 = id
`

func (q *Queries) FindByID(ctx context.Context, id int32) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const findByIDAndName = `-- name: FindByIDAndName :many
SELECT id, name FROM users WHERE $1 = id AND $1 = name
`

func (q *Queries) FindByIDAndName(ctx context.Context, id int32) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findByIDAndName, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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
