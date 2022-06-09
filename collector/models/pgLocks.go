package models

import (
	"database/sql"
)

// +metric=slice
type PgLocks struct {
	tableName struct{}       `pg:"pg_locks"`
	Locktype  sql.NullString `pg:"locktype" help:"Type of the lockable object" metric:",type:label"`
	ScopeType sql.NullString `pg:"scope_type" help:"The type of the target" metric:",type:label"`
	Database  sql.NullString `pg:"database" help:"The Database name if applicable" metric:",type:label"`
	Mode      sql.NullString `pg:"mode" help:"Name of the lock mode held or desired by this process" metric:",type:label"`
	Granted   sql.NullBool   `pg:"granted" help:"True if lock is held, false if lock is awaited" metric:",type:label"`
	Locks     int64          `pg:"locks" help:"Number of locks per state" metric:"count,type:gauge"`
}
