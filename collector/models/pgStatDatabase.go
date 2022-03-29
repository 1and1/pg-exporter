package models

import (
	"time"
)

// +metric=slice
type PgStatDatabase struct {
	tableName    struct{}     `pg:"pg_stat_database"`
	DatID        int64        `pg:"datid" help:"OID of a database" metric:"database_id,type:label"`
	DatName      string       `pg:"datname" help:"Name of this database" metric:"database,type:label"`
	NumBackends  int          `pg:"numbackends" help:"Number of backends currently connected to this database" metric:"backends,type:gauge"`
	XactCommit   int64        `pg:"xact_commit" help:"Number of transactions in this database that have been committed" metric:"xact_commited_total"`
	XactRollback int64        `pg:"xact_rollback" help:"Number of transactions in this database that have been rolled back" metric:"xact_rolledback_total"`
	BlksRead     int64        `pg:"blks_read" help:"Number of disk blocks read in this database" metric:"blocks_read_total"`
	BlksHit      int64        `pg:"blks_hit" help:"Number of times disk blocks were found already in the buffer cache, so that a read was not necessary" metric:"blocks_hit_total"`
	TupReturned  int64        `pg:"tup_returned" help:"Number of rows returned by queries in this database" metric:"rows_returned_total"`
	TupFetched   int64        `pg:"tup_fetched" help:"Number of rows fetched by queries in this database" metric:"rows_fetched_total"`
	TupInserted  int64        `pg:"tup_inserted" help:"Number of rows inserted by queries in this database" metric:"rows_inserted_total"`
	TupUpdated   int64        `pg:"tup_updated" help:"Number of rows updated by queries in this database" metric:"rows_updated_total"`
	TupDeleted   int64        `pg:"tup_deleted" help:"Number of rows deleted by queries in this database" metric:"rows_deleted_total"`
	Conflicts    int64        `pg:"conflicts" help:"Number of queries canceled due to conflicts with recovery in this database"  metric:"conflicts_total"`
	TempFiles    int64        `pg:"temp_files" help:"Number of temporary files created by queries in this database"  metric:"temp_files_total"`
	TempBytes    int64        `pg:"temp_bytes" help:"Total amount of data written to temporary files by queries in this database"  metric:"temp_bytes_total"`
	Deadlocks    int64        `pg:"deadlocks" help:"Number of deadlocks detected in this database"  metric:"deadlocks_total"`
	ChecksumFailures int64    `pg:"checksum_failures" help:"Number of data page checksum failures detected in this database" metric:"checksum_failures_count"` // new in PG12
	ChecksumLastFailure time.Time `pg:"checksum_last_failure" help:"Time at which the last data page checksum failure was detected in this database" metric:"checksum_last_failure"` // new in PG12
	BlkReadTime  Milliseconds `pg:"blk_read_time" help:"Time spent reading data file blocks by backends in this database" metric:"blk_read_seconds_total"`
	BlkWriteTime Milliseconds `pg:"blk_write_time" help:"Time spent writing data file blocks by backends in this database" metric:"blk_write_seconds_total"`
	SessionTime  Milliseconds `pg:"session_time" help:"Time spent by database sessions in this database, in milliseconds" metric:"session_time_total"` // new in PG14
	ActiveTime   Milliseconds `pg:"active_time" help:"Time spent executing SQL statements in this database, in milliseconds" metric:"active_time_total"` // new in PG14
	IdleInTransactionTime Milliseconds `pg:"idle_in_transaction_time" help:"Time spent idling while in a transaction in this database, in milliseconds" metric:"idle_in_transaction_time_total"` // new in PG14
	Sessions     int64        `pg:"sessions" help:"Total number of sessions established to this database" metric:"sessions_count"` // new in PG14
	SessionsAbandoned int64   `pg:"sessions_abandoned" help:"Number of database sessions to this database that were terminated because connection to the client was lost" metric:"sessions_abandoned_count"` // new in PG14
	SessionsFatal int64       `pg:"sessions_fatal" help:"Number of database sessions to this database that were terminated by fatal errors" metric:"sessions_fatal_count"` // new in PG14
	SessionsKilled int64      `pg:"sessions_killed" help:"Number of database sessions to this database that were terminated by operator intervention" metric:"sessions_killed_count"` // new in PG14
	StatsReset   time.Time    `pg:"stats_reset" help:"Time at which these statistics were last reset"`
}
