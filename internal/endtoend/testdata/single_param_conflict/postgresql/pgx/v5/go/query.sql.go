// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAuthorByID = `-- name: GetAuthorByID :one
SELECT  id, name, bio
FROM    authors
WHERE   id = $1
LIMIT   1
`

func (q *Queries) GetAuthorByID(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthorByID, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getAuthorIDByID = `-- name: GetAuthorIDByID :one
SELECT  id
FROM    authors
WHERE   id = $1
LIMIT   1
`

func (q *Queries) GetAuthorIDByID(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, getAuthorIDByID, id)
	err := row.Scan(&id)
	return id, err
}

const getUser = `-- name: GetUser :one
SELECT  sub
FROM    users
WHERE   sub = $1
LIMIT   1
`

func (q *Queries) GetUser(ctx context.Context, sub pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, getUser, sub)
	err := row.Scan(&sub)
	return sub, err
}

const setDefaultName = `-- name: SetDefaultName :one

UPDATE  authors
SET     name = 'Default Name'
WHERE   id = $1
RETURNING id
`

// https://github.com/sqlc-dev/sqlc/issues/1235
func (q *Queries) SetDefaultName(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, setDefaultName, id)
	err := row.Scan(&id)
	return id, err
}
