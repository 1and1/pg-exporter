package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

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
func (ScrapePreparedXacts) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	qs := `SELECT count(1) AS count, min(prepared) AS oldest, database
            FROM pg_prepared_xacts
            WHERE database IN (?)
            GROUP BY database;`
	var xacts models.PgPreparedTransactionsSlice
	if _, err := db.QueryContext(ctx, &xacts, qs, pg.In(collectDatabases)); err != nil {
		return err
	}
	return xacts.ToMetrics(namespace, preparedXacts, ch)
}
