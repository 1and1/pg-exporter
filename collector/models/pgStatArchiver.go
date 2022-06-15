package models

import (
	"time"

	"github.com/uptrace/bun"
)

// +metric=row
type PgStatArchiver struct {
	bun.BaseModel    `bun:"pg_stat_archiver"`
	ArchivedCount    int64     `bun:"archived_count" help:"Number of WAL files that have been successfully archived" metric:"archived_count_total"`
	LastArchivedTime time.Time `bun:"last_archived_time" help:"Time of the last successful archive operation"`
	FailedCount      int64     `bun:"failed_count" help:"Number of failed attempts for archiving WAL files"  metric:"failed_count_total"`
	LastFailedTime   time.Time `bun:"last_failed_time" help:"Time of the last failed archival operation"`
	StatsReset       time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset"`
}
