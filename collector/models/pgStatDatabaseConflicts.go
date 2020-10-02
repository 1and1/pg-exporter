package models

// +metric=slice
type PgStatDatabaseConflicts struct {
	tableName      struct{} `sql:"pg_stat_database_conflicts"`
	DatID          int64    `sql:"datid" help:"OID of a database" metric:"database_id,type:label"`
	DatName        string   `sql:"datname" help:"Name of this database" metric:"database,type:label"`
	ConfTablespace int64    `sql:"confl_tablespace" help:"Number of queries in this database that have been canceled due to dropped tablespaces" metric:"confl_tablespace_total"`
	ConfLock       int64    `sql:"confl_lock" help:"Number of queries in this database that have been canceled due to lock timeouts" metric:"confl_lock_total"`
	ConfSnapshot   int64    `sql:"confl_snapshot" help:"Number of queries in this database that have been canceled due to old snapshots" metric:"confl_snapshot_total"`
	ConfBufferpin  int64    `sql:"confl_bufferpin" help:"Number of queries in this database that have been canceled due to pinned buffers" metric:"confl_bufferpin_total"`
	ConfDeadlock   int64    `sql:"confl_deadlock" help:"Number of queries in this database that have been canceled due to deadlocks" metric:"confl_deadlock_total"`
}
