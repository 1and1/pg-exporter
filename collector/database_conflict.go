package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	databaseconflicts = "database_conflicts"
)

// ScrapeBgWriter scrapes from pg_stat_database_conflicts
type ScrapeDatabaseConflicts struct{}

// Name of the Scraper
func (ScrapeDatabaseConflicts) Name() string {
	return "pg_stat_database_conflicts"
}

// Help describes the role of the Scraper
func (ScrapeDatabaseConflicts) Help() string {
	return "Collect from pg_stat_database_conflicts"
}

// minimum PostgreSQL version
func (ScrapeDatabaseConflicts) Version() int {
	return 0
}

// scrape type
func (ScrapeDatabaseConflicts) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeDatabaseConflicts) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	var databaseConflict models.PgStatDatabaseConflictsSlice
	if err := db.ModelContext(ctx, &databaseConflict).Where("datname IN (?)", pg.In(collectDatabases)).
		Select(); err != nil {
		return err
	}

	return databaseConflict.ToMetrics(namespace, databaseconflicts, ch)
}
