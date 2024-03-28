// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: copyfrom_gen.go

package querytest

import (
	"context"
)

// iteratorForUsersC implements pgx.CopyFromSource.
type iteratorForUsersC struct {
	rows                 []int64
	skippedFirstNextCall bool
}

func (r *iteratorForUsersC) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForUsersC) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0],
	}, nil
}

func (r iteratorForUsersC) Err() error {
	return nil
}

func (q *Queries) UsersC(ctx context.Context, id []int64) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"user"}, []string{"id"}, &iteratorForUsersC{rows: id})
}
