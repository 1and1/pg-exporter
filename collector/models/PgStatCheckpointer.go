package models

import (
    "time"

    "github.com/uptrace/bun"
)

// +metric=row
type PgStatCheckpointer struct {
    bun.BaseModel      `bun:"pg_stat_checkpointer"`
    NumTimed           int64     `bun:"num_timed" help:"Number of Checkpoints  timed" metric:"num_timed_total"`
    NumRequested       int64     `bun:"num_requested" help:"Number of Checkpoints requested" metric:"num_requested_total"`
    RestartpointsTimed int64     `bun:"restartpoints_timed" help:"Number of Restart Points timed" metric:"restartpoints_timed_total"`
    RestartpointsReq   int64     `bun:"restartpoints_req" help:"Number of Restart Points requested" metric:"restartpoints_req_total"`
    RestartpointsDone  int64     `bun:"restartpoints_done" help:"Number of Restart Points done" metric:"restartpoints_done_total"`
    CkpWriteTime       float64   `bun:"write_time" help:"Checkpoint write Time" metric:"ckp_write_time_total"`
    CkpSyncTime        float64   `bun:"sync_time" help:"Checkpoint sync Time" metric:"ckp_sync_time_total"`
    BuffersWritten     int64     `bun:"buffers_written"  help:"Checkpoint Buffers written" metric:"ckp_buffers_written_total"`
    StatsReset         time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset"`
}
