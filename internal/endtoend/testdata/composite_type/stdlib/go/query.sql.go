// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const listPaths = `-- name: ListPaths :many
SELECT point_one, point_two FROM foo.paths
`

func (q *Queries) ListPaths(ctx context.Context) ([]FooPath, error) {
	rows, err := q.db.QueryContext(ctx, listPaths)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FooPath
	for rows.Next() {
		var i FooPath
		if err := rows.Scan(&i.PointOne, &i.PointTwo); err != nil {
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
