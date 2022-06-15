package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	iotables = "tablesio"
)

// ScrapeIOTables scrapes from pg_statio_user_tables
type ScrapeIOTables struct{}

// Name of the Scraper
func (ScrapeIOTables) Name() string {
	return "pg_statio_user_tables"
}

// Help describes the role of the Scraper
func (ScrapeIOTables) Help() string {
	return "Collect from pg_statio_user_tables"
}

// Version returns minimum PostgreSQL version
func (ScrapeIOTables) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeIOTables) Type() ScrapeType {
	return SCRAPELOCAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeIOTables) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	// get the db name we are connected to, we need it as a label
	row, err := db.QueryContext(ctx, "SELECT current_database()")
	if err != nil {
		return err
	}
	var dbname string
	if err := db.ScanRows(ctx, row, &dbname); err != nil {
		return err
	}
	row.Close()

	var statIOUserTables models.PgStatIOUserTablesSlice
	if err := db.NewSelect().Model(&statIOUserTables).Where("schemaname NOT LIKE ?", "pg_temp_%").Scan(ctx); err != nil {
		return err
	}

	return statIOUserTables.ToMetrics(namespace, iotables, ch, "database", dbname)
}
