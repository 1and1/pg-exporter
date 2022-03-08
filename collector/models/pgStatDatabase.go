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
	BlkReadTime  Milliseconds `pg:"blk_read_time" help:"Time spent reading data file blocks by backends in this database" metric:"blk_read_seconds_total"`
	BlkWriteTime Milliseconds `pg:"blk_write_time" help:"Time spent writing data file blocks by backends in this database" metric:"blk_write_seconds_total"`
	StatsReset   time.Time    `pg:"stats_reset" help:"Time at which these statistics were last reset"`
}
