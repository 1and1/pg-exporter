package models

import (
	"database/sql"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgLocks struct {
	bun.BaseModel `bun:"pg_locks"`
	Locktype      sql.NullString `bun:"locktype" help:"Type of the lockable object" metric:",type:label"`
	ScopeType     sql.NullString `bun:"scope_type" help:"The type of the target" metric:",type:label"`
	Database      sql.NullString `bun:"database" help:"The Database name if applicable" metric:",type:label"`
	Mode          sql.NullString `bun:"mode" help:"Name of the lock mode held or desired by this process" metric:",type:label"`
	Granted       sql.NullBool   `bun:"granted" help:"True if lock is held, false if lock is awaited" metric:",type:label"`
	Locks         int64          `bun:"locks" help:"Number of locks per state" metric:"count,type:gauge"`
}
