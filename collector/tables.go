package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	tables = "tables"
)

// ScrapeTables scrapes from pg_stat_user_tables
type ScrapeTables struct{}

// Name of the Scraper
func (ScrapeTables) Name() string {
	return "pg_stat_user_tables"
}

// Help describes the role of the Scraper
func (ScrapeTables) Help() string {
	return "Collect from pg_stat_user_tables"
}

// Version returns minimum PostgreSQL version
func (ScrapeTables) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeTables) Type() ScrapeType {
	return SCRAPELOCAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeTables) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
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

	var statUserTables models.PgStatUserTablesSlice
	if err := db.NewSelect().Model(&statUserTables).Where("schemaname NOT LIKE ?", "pg_temp_%").Scan(ctx); err != nil {
		return err
	}

	return statUserTables.ToMetrics(namespace, tables, ch, "database", dbname)
}
