package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	wal = "wal"
)

// ScrapeWal scrapes from pg_stat_wal
type ScrapeWal struct{}

// Name of the Scraper
func (ScrapeWal) Name() string {
	return "pg_stat_wal"
}

// Help describes the role of the Scraper
func (ScrapeWal) Help() string {
	return "Collect from pg_stat_wal"
}

// Version returns minimum PostgreSQL version
func (ScrapeWal) Version() int {
	return 140000
}

// Type returns the scrape type
func (ScrapeWal) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeWal) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	statWal := &models.PgStatWal{}
	if err := db.NewSelect().Model(statWal).Scan(ctx); err != nil {
		return err
	}

	return statWal.ToMetrics(namespace, wal, ch)
}
