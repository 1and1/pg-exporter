package models

import (
	"time"
)

// +metric=row
type PgStatArchiver struct {
	tableName        struct{}   `sql:"pg_stat_archiver"`
	ArchivedCount    int64      `sql:"archived_count" help:"Number of WAL files that have been successfully archived" metric:"archived_count_total"`
	LastArchivedTime time.Time `sql:"last_archived_time" help:"Time of the last successful archive operation"`
	FailedCount      int64      `sql:"failed_count" help:"Number of failed attempts for archiving WAL files"  metric:"failed_count_total"`
	LastFailedTime   time.Time `sql:"last_failed_time" help:"Time of the last failed archival operation"`
	StatsReset       time.Time `sql:"stats_reset" help:"Time at which these statistics were last reset"`
}
