package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	tables = "tables"
)

// ScrapeBgWriter scrapes from pg_stat_user_tables
type ScrapeTables struct{}

// Name of the Scraper
func (ScrapeTables) Name() string {
	return "pg_stat_user_tables"
}

// Help describes the role of the Scraper
func (ScrapeTables) Help() string {
	return "Collect from pg_stat_user_tables"
}

// minimum PostgreSQL version
func (ScrapeTables) Version() int {
	return 0
}

// scrape type
func (ScrapeTables) Type() ScrapeType {
	return SCRAPELOCAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeTables) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	// get the db name we are connected to, we need it as a label
	dbname := db.Options().Database

	var statUserTables models.PgStatUserTablesSlice
	if err := db.ModelContext(ctx, &statUserTables).Where("schemaname NOT LIKE ?", "pg_temp_%").Select(); err != nil {
		return err
	}

	return statUserTables.ToMetrics(namespace, tables, ch, "database", dbname)
}
