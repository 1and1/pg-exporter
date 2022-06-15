package models

import (
	"database/sql"
)

// +metric=row
type PgTxid struct {
	Current sql.NullInt64 `bun:"current" help:"The current transaction ID on the database cluster"`
}
