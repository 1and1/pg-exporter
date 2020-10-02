package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	iotables = "tablesio"
)

// ScrapeBgWriter scrapes from pg_stat_user_tables
type ScrapeIOTables struct{}

// Name of the Scraper
func (ScrapeIOTables) Name() string {
	return "pg_statio_user_tables"
}

// Help describes the role of the Scraper
func (ScrapeIOTables) Help() string {
	return "Collect from pg_statio_user_tables"
}

// minimum PostgreSQL version
func (ScrapeIOTables) Version() int {
	return 0
}

// scrape type
func (ScrapeIOTables) Type() ScrapeType {
	return SCRAPELOCAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeIOTables) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	// get the db name we are connected to, we need it as a label
	dbname := db.Options().Database

	var statIOUserTables models.PgStatIOUserTablesSlice
	if err := db.ModelContext(ctx, &statIOUserTables).Where("schemaname NOT LIKE ?", "pg_temp_%").Select(); err != nil {
		return err
	}

	return statIOUserTables.ToMetrics(namespace, iotables, ch, "database", dbname)
}
