package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	txid = "txid"
)

// ScrapeTXID scrapes from txid_current()
type ScrapeTXID struct{}

// Name of the Scraper
func (ScrapeTXID) Name() string {
	return "txid_current"
}

// Help describes the role of the Scraper
func (ScrapeTXID) Help() string {
	return "Collect from txid_current()"
}

// Version returns minimum PostgreSQL version
func (ScrapeTXID) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeTXID) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeTXID) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	qs := `SELECT CASE WHEN pg_is_in_recovery() THEN NULL ELSE txid_current() END AS current`
	ctxid := &models.PgTxid{}
	rows, err := db.QueryContext(ctx, qs)
	if err != nil {
		return err
	}
	if err := db.ScanRows(ctx, rows, ctxid); err != nil {
		return err
	}

	return ctxid.ToMetrics(namespace, txid, ch)
}
