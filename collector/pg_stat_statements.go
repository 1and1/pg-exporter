package collector

import (
"context"

"github.com/go-pg/pg/v9"
"github.com/prometheus/client_golang/prometheus"

"github.com/1and1/pg-exporter/collector/models"
)

const (
    // subsystem
    statstatements = "statements"
)

// ScrapeStatements scrapes from pg_stat_statements
type ScrapeStatements struct{}

// Name of the Scraper
func (ScrapeStatements) Name() string {
    return "pg_stat_statements"
}

// Help describes the role of the Scraper
func (ScrapeStatements) Help() string {
    return "Collect from pg_stat_statements"
}

// Version returns minimum PostgreSQL version
func (ScrapeStatements) Version() int {
    return 0
}

// Type returns the scrape type
func (ScrapeStatements) Type() ScrapeType {
    return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeStatements) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
    var statDatabase models.PgStatDatabaseSlice
    if err := db.ModelContext(ctx, &statDatabase).Where("datname IN (?)", pg.In(collectDatabases)).
        Select(); err != nil {
        return err
    }

    return statDatabase.ToMetrics(namespace, statstatements, ch)
}

