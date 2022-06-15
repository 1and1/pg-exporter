package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	archiver = "archiver"
)

// ScrapeArchiver scrapes from pg_stat_archiver
type ScrapeArchiver struct{}

// Name of the Scraper
func (ScrapeArchiver) Name() string {
	return "pg_stat_archiver"
}

// Help describes the role of the Scraper
func (ScrapeArchiver) Help() string {
	return "Collect from pg_stat_archiver"
}

// Version returns minimum PostgreSQL version
func (ScrapeArchiver) Version() int {
	return 0
}

// Type returns the scrape type
func (ScrapeArchiver) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeArchiver) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	statArchiver := &models.PgStatArchiver{}
	if err := db.NewSelect().Model(statArchiver).Scan(ctx); err != nil {
		return err
	}

	return statArchiver.ToMetrics(namespace, archiver, ch)
}
