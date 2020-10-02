package collector

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/prometheus/client_golang/prometheus"

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

// minimum PostgreSQL version
func (ScrapeTXID) Version() int {
	return 0
}

// scrape type
func (ScrapeTXID) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeTXID) Scrape(ctx context.Context, db *pg.DB, ch chan<- prometheus.Metric) error {
	qs := `SELECT CASE WHEN pg_is_in_recovery() THEN NULL ELSE txid_current() END AS current`
	ctxid := &models.PgTxid{}
	if _, err := db.QueryContext(ctx, ctxid, qs); err != nil {
		return err
	}
	return ctxid.ToMetrics(namespace, txid, ch)
}
