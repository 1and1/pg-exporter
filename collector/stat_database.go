package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

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
func (ScrapeDatabase) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	var statDatabase models.PgStatDatabaseSlice
	if err := db.ModelContext(ctx, &statDatabase).Where("datname IN (?)", pg.In(collectDatabases)).
		Select(); err != nil {
		return err
	}

	if err := statDatabase.ToMetrics(namespace, statdatabase, ch); err != nil {
		return err
	}

	var databases models.PgDatabaseSlice
	if err := db.ModelContext(ctx, &databases).Where("datname IN (?)", pg.In(collectDatabases)).
		Select(); err != nil {
		return err
	}
	return databases.ToMetrics(namespace, statdatabase, ch)
}
