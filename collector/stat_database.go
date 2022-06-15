package collector

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	statdatabase = "database"
)

// ScrapeDatabase scrapes from pg_stat_database
type ScrapeDatabase struct{}

// Name of the Scraper
func (ScrapeDatabase) Name() string {
	return "pg_stat_database"
}

// Help describes the role of the Scraper
func (ScrapeDatabase) Help() string {
	return "Collect from pg_stat_database"
}

// Version returns minimum PostgreSQL version
func (ScrapeDatabase) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeDatabase) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeDatabase) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	// we create a query based on the given commandline flags
	columns := "numbackends, xact_commit, xact_rollback, blks_read, blks_hit, tup_returned, tup_fetched, tup_inserted, tup_updated, tup_deleted, conflicts, temp_files, temp_bytes, deadlocks, blk_read_time, blk_write_time"

	if pgversion >= 120000 {
		columns += ", checksum_failures, checksum_last_failure"
	}

	if pgversion >= 140000 {
		columns += ", session_time, active_time, idle_in_transaction_time, sessions, sessions_abandoned, sessions_fatal, sessions_killed"
	}

	qs := fmt.Sprintf(`SELECT datid, datname, %s, stats_reset FROM pg_stat_database`+
		` WHERE datname IN (?)`,
		columns)

	var statDatabase models.PgStatDatabaseSlice
	rows, err := db.QueryContext(ctx, qs, bun.In(collectDatabases))
	if err != nil {
		return err
	}
	if err := db.ScanRows(ctx, rows, &statDatabase); err != nil {
		return err
	}
	if err := statDatabase.ToMetrics(namespace, statdatabase, ch); err != nil {
		return err
	}

	var databases models.PgDatabaseSlice
	if err := db.NewSelect().Model(&databases).Where("datname IN (?)", bun.In(collectDatabases)).
		Scan(ctx); err != nil {
		return err
	}
	return databases.ToMetrics(namespace, statdatabase, ch)
}
