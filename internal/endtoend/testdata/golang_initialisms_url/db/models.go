// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Foo struct {
	BarID   sql.NullString
	SiteURL sql.NullString
}
