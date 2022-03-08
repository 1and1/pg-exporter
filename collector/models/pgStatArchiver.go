package models

import (
	"time"
)

// +metric=row
type PgStatArchiver struct {
	tableName        struct{}   `pg:"pg_stat_archiver"`
	ArchivedCount    int64      `pg:"archived_count" help:"Number of WAL files that have been successfully archived" metric:"archived_count_total"`
	LastArchivedTime time.Time `pg:"last_archived_time" help:"Time of the last successful archive operation"`
	FailedCount      int64      `pg:"failed_count" help:"Number of failed attempts for archiving WAL files"  metric:"failed_count_total"`
	LastFailedTime   time.Time `pg:"last_failed_time" help:"Time of the last failed archival operation"`
	StatsReset       time.Time `pg:"stats_reset" help:"Time at which these statistics were last reset"`
}
