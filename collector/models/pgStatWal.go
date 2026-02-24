package models

import (
	"time"
    "database/sql"
	"github.com/uptrace/bun"
)

// +metric=row
type PgStatWal struct {
	bun.BaseModel  `bun:"pg_stat_wal"`
	WalRecords     int64     `bun:"wal_records" help:"Total number of WAL records generated" metric:"records_total"`
	WalFpi         int64     `bun:"wal_fpi" help:"Total number of WAL full page images generated" metric:"fpi_total"`
	WalBytes       float64   `bun:"wal_bytes" help:"Total amount of WAL generated in bytes" metric:"bytes_total"`
	WalBuffersFull int64     `bun:"wal_buffers_full" help:"Number of times WAL data was written to disk because WAL buffers became full" metric:"buffers_full_count"`
	WalWrite       sql.NullInt64     `bun:"wal_write" help:"Number of times WAL buffers were written out to disk via XLogWrite request" metric:"write_count"`
	WalSync        sql.NullInt64     `bun:"wal_sync" help:"Number of times WAL files were synced to disk via issue_xlog_fsync request" metric:"sync_count"`
	WalWriteTime   sql.NullFloat64   `bun:"wal_write_time" help:"Total amount of time spent writing WAL buffers to disk via XLogWrite request, in milliseconds" metric:"write_time_total"`
	WalSyncTime    sql.NullFloat64   `bun:"wal_sync_time" help:"Total amount of time spent syncing WAL files to disk via issue_xlog_fsync request, in milliseconds" metric:"sync_time_total"`
	StatsReset     time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset"`
}
