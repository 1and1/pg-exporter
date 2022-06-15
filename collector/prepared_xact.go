package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	preparedXacts = "prepared_transactions"
)

// ScrapePreparedXacts scrapes from pg_prepared_xacts
type ScrapePreparedXacts struct{}

// Name of the Scraper
func (ScrapePreparedXacts) Name() string {
	return "pg_prepared_xacts"
}

// Help describes the role of the Scraper
func (ScrapePreparedXacts) Help() string {
	return "Collect from pg_prepared_xacts"
}

// Version returns minimum PostgreSQL version
func (ScrapePreparedXacts) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapePreparedXacts) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapePreparedXacts) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	qs := `SELECT count(1) AS count, min(prepared) AS oldest, database
            FROM pg_prepared_xacts
            WHERE database IN (?)
            GROUP BY database;`
	var xacts models.PgPreparedTransactionsSlice
	rows, err := db.QueryContext(ctx, qs, bun.In(collectDatabases))
	if err != nil {
		return err
	}
	if err := db.ScanRows(ctx, rows, &xacts); err != nil {
		return err
	}
	return xacts.ToMetrics(namespace, preparedXacts, ch)
}
