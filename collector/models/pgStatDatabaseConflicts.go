package models

import (
	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatDatabaseConflicts struct {
	bun.BaseModel  `bun:"pg_stat_database_conflicts"`
	DatID          int64  `bun:"datid" help:"OID of a database" metric:"database_id,type:label"`
	DatName        string `bun:"datname" help:"Name of this database" metric:"database,type:label"`
	ConfTablespace int64  `bun:"confl_tablespace" help:"Number of queries in this database that have been canceled due to dropped tablespaces" metric:"confl_tablespace_total"`
	ConfLock       int64  `bun:"confl_lock" help:"Number of queries in this database that have been canceled due to lock timeouts" metric:"confl_lock_total"`
	ConfSnapshot   int64  `bun:"confl_snapshot" help:"Number of queries in this database that have been canceled due to old snapshots" metric:"confl_snapshot_total"`
	ConfBufferpin  int64  `bun:"confl_bufferpin" help:"Number of queries in this database that have been canceled due to pinned buffers" metric:"confl_bufferpin_total"`
	ConfDeadlock   int64  `bun:"confl_deadlock" help:"Number of queries in this database that have been canceled due to deadlocks" metric:"confl_deadlock_total"`
}
