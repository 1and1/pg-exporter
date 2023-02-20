package collector

import (
	"context"

	"github.com/1and1/pg-exporter/collector/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"
)

const (
	// subsystem
	frozenxid = "frozenxid"
)

// ScrapeFrozenXid scrapes from pg_stat_user_tables
type ScrapeFrozenXid struct{}

// Name of the Scraper
func (ScrapeFrozenXid) Name() string {
	return "relfrozenxid"
}

// Help describes the role of the Scraper
func (ScrapeFrozenXid) Help() string {
	return "Collect relfrozenxid pg_class"
}

// Version returns minimum PostgreSQL version
func (ScrapeFrozenXid) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeFrozenXid) Type() ScrapeType {
	return SCRAPELOCAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeFrozenXid) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	row, err := db.QueryContext(ctx, "SELECT current_database()")
	if err != nil {
		return err
	}
	var dbname string
	if err := db.ScanRows(ctx, row, &dbname); err != nil {
		return err
	}
	row.Close()

	var frozenxids models.PgFrozenXidSlice

	err = db.NewRaw(`SELECT c.relnamespace::regnamespace AS schema,
       c.oid::regclass              AS table_name,
       AGE(c.relfrozenxid)          AS relfrozenxid,
       AGE(t.relfrozenxid)          AS toastfrozenxid
FROM pg_class c
         LEFT JOIN pg_class t ON c.reltoastrelid = t.oid
WHERE c.relkind IN ('r', 'm')
  AND c.relnamespace::regnamespace NOT IN ('pg_catalog', 'information_schema')`).Scan(ctx, &frozenxids)

	// get the db name we are connected to, we need it as a label
	if err != nil {
		return err
	}

	return frozenxids.ToMetrics(namespace, frozenxid, ch, "database", dbname)
}
