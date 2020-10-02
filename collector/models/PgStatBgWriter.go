package models

import (
	"time"
)

// +metric=row
type PgStatBgWriter struct {
	tableName           struct{}  `sql:"pg_stat_bgwriter"`
	CheckpointsTimed    int64     `sql:"checkpoints_timed" help:"Number of scheduled checkpoints that have been performed" metric:"checkpoints_timed_total"`
	CheckpointsReq      int64     `sql:"checkpoints_req" help:"Number of requested checkpoints that have been performed" metric:"checkpoints_req_total"`
	CheckpointWriteTime float64   `sql:"checkpoint_write_time" help:"Total amount of time that has been spent in the portion of checkpoint processing where files are written to disk, in milliseconds" metric:"checkpoint_write_seconds_total"`
	CheckpointSyncTime  float64   `sql:"checkpoint_sync_time" help:"Total amount of time that has been spent in the portion of checkpoint processing where files are synchronized to disk, in milliseconds"  metric:"checkpoint_sync_seconds_total"`
	BuffersCheckpoint   int64     `sql:"buffers_checkpoint" help:"Number of buffers written during checkpoints" metric:"buffers_checkpoint_total"`
	BuffersClean        int64     `sql:"buffers_clean" help:"Number of buffers written by the background writer" metric:"buffers_clean_total"`
	MaxwrittenClean     int64     `sql:"maxwritten_clean" help:"Number of times the background writer stopped a cleaning scan because it had written too many buffers" metric:"maxwritten_clean_total"`
	BuffersBackend      int64     `sql:"buffers_backend" help:"Number of buffers written directly by a backend" metric:"buffers_backend_total"`
	BuffersBackendFsync int64     `sql:"buffers_backend_fsync" help:"Number of times a backend had to execute its own fsync call" metric:"buffers_backend_fsync_total"`
	BuffersAlloc        int64     `sql:"buffers_alloc" help:"Number of buffers allocated" metric:"buffers_alloc_total"`
	StatsReset          time.Time `sql:"stats_reset" help:"Time at which these statistics were last reset"`
}
