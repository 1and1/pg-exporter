package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	locks = "locks"
)

// ScrapeLocks scrapes from pg_locks
type ScrapeLocks struct{}

// Name of the Scraper
func (ScrapeLocks) Name() string {
	return "pg_locks"
}

// Help describes the role of the Scraper
func (ScrapeLocks) Help() string {
	return "Collect from pg_locks"
}

// Version returns minimum PostgreSQL version
func (ScrapeLocks) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeLocks) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeLocks) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	qs := `WITH locks AS (
            SELECT l.locktype,
                   CASE
                       WHEN l.database IS NULL THEN 'transaction'
                       WHEN l.database = 0 THEN 'shared object'
                       ELSE 'database'
                       END AS scope_type,
                   CASE
                       WHEN l.database IS NOT NULL AND l.database > 0
                           THEN (SELECT datname FROM pg_database db WHERE db.oid = l.database)
                       END AS database,
                   l.mode,
                   l.granted
            FROM pg_locks l
        )
        SELECT locktype, scope_type, database, "mode", "granted", count(1) AS locks
        FROM locks
        WHERE database IS NULL OR database IN (?)
        GROUP BY locktype, scope_type, database, "mode", "granted";`

	var dblocks models.PgLocksSlice
	if _, err := db.QueryContext(ctx, &dblocks, qs, pg.In(collectDatabases)); err != nil {
		return err
	}
	return dblocks.ToMetrics(namespace, locks, ch)
}
