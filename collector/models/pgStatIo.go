package models

import (
	"time"
	"github.com/uptrace/bun"
    "database/sql"
)

// +metric=row
type PgStatIo struct {
	bun.BaseModel  `bun:"pg_stat_io"`
	BackendType    string    `bun:"backend_type" help:"Type of backend (e.g. background worker, autovacuum worker)" metric:",type:label"`
    Object         string    `bun:"object" help:"Target object of an I/O operation (e.g. relation, temp relation, wal)" metric:",type:label"`
	Context        string    `bun:"context" help:"The context of an I/O operation" metric:",type:label"`
    OpBytes        sql.NullInt64     `bun:"op_bytes" help:"The number of bytes per unit of I/O read, written, or extended" metric:"statio_op_bytes_total"`
    Reads          int64     `bun:"reads" help:"Number of read operations" metric:"statio_reads_total"`
    ReadBytes      sql.NullFloat64   `bun:"read_bytes" help:"The total size of read operations in bytes" metric:"statio_read_bytes_total"`
    ReadTime       float64   `bun:"read_time" help:"Time spent waiting for read operations in milliseconds" metric:"statio_read_time"`
    Writes         int64     `bun:"writes" help:"Number of write operations"  metric:"statio_writes_total"`
    WriteBytes     sql.NullFloat64 `bin:"write_bytes" help:"The total size of write operations in bytes"  metric:"statio_write_bytes_total"`
    WriteTimes     float64   `bun:"write_time" help:"Time spent waiting for write operations in milliseconds"  metric:"statio_write_time"`
    WriteBacks     int64     `bun:"writebacks" help:"Number of units of size BLCKSZ (typically 8kB) which the process requested the kernel write out to permanent storage" metric:"statio_writebacks_total"`
    WriteBackTime  float64   `bun:"writeback_time" help:"Time spent waiting for writeback operations in milliseconds (if track_io_timing is enabled, otherwise zero). This includes the time spent queueing write-out requests and, potentially, the time spent to write out the dirty data"  metric:"statio_writeback_time"`
    Extends        int64     `bun:"extends" help:"Number of relation extend operations"  metric:"statio_extends_total"`
    ExtendBytes    sql.NullFloat64   `bin:"extend_bytes" help:"The total size of relation extend operations in bytes" metric:"statio_extend_bytes_total"`
    ExtendTime     float64   `bun:"extend_time" help:"Time spent waiting for extend operations in milliseconds. (if track_io_timing is enabled and object is not wal, or if track_wal_io_timing is enabled and object is wal, otherwise zero)" metric:"statio_extend_time"`
    Hits           int64     `bun:"hits"  help:"The number of times a desired block was found in a shared buffer" metric:"statio_hits_total"`
    Evictions      int64     `bun:"evictions"  help:"Number of times a block has been written out from a shared or local buffer in order to make it available for another use"  metrics:"statio_evictions_total"`
    Reuses         int64     `bun:"reuses"  help:"The number of times an existing buffer in a size-limited ring buffer outside of shared buffers was reused as part of an I/O operation in the bulkread, bulkwrite, or vacuum contexts" metric:"statio_reuses_total"`
    Fsyncs         int64     `bun:"fsyncs"  help:"Number of fsync calls. These are only tracked in context normal"  metric:"statio_fsyncs_total"`
    FsyncTime      float64   `bun:"fsync_time"  help:"Time spent waiting for fsync operations in milliseconds" metric:"statio_fsync_time"`
	StatsReset     time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset"`
}
