// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package querytest

import (
	"database/sql"
)

type Foo struct {
	Bar    string
	Baz    sql.NullInt32
	Bio    sql.NullInt32
	Foobar []int32
}
